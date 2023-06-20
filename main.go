package main

import (
	"context"
	"gitbot/controllers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"gitbot/configs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := configs.GetConfig()

	// Router
	path := "/" + cfg.PathURL + "/{id}"
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post(path, controllers.HandleWebHook)

	// server config
	srv := &http.Server{
		Addr:        ":" + cfg.Port,
		Handler:     r,
		ReadTimeout: 10 * time.Second,
	}

	// database
	if _, err := controllers.LoadDatabase(); err != nil {
		log.Fatalln(err)
	}

	// Listening to interrupt signal
	var wg sync.WaitGroup
	wg.Add(1)

	idleConnsClosed := make(chan struct{})
	go func() {
		defer wg.Done()
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint

		log.Printf("service interrupt received\n")
		log.Printf("http server shutting down\n")

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		if err := controllers.CloseDatabase(ctx); err != nil {
			log.Fatalln(err)
		}
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("http server shutdown error: %v", err)
		}

		log.Printf("shutdown completed\n")
		close(idleConnsClosed)
	}()

	// Handle Telegram Command
	go controllers.HandleCommand()

	// Serve
	log.Printf("Listening to port %s.\n", cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		if err.Error() != "http: Server closed" {
			log.Printf("HTTP server closed with: %v\n", err)
		}
	}

	wg.Wait()
}
