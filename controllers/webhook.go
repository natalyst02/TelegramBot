package controllers

import (
	"context"
	"encoding/json"
	"gitbot/configs"
	"gitbot/models"
	"gitbot/models/webhook"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleWebHook(w http.ResponseWriter, r *http.Request) {
	chiParam := chi.URLParam(r, "id")

	var findRes bson.M
	uid, err := primitive.ObjectIDFromHex(chiParam)
	if err != nil {
		w.WriteHeader(500)
		log.Panic(err)
	}

	err = GetCol().FindOne(context.TODO(), bson.D{{Key: "_id", Value: uid}}).Decode(&findRes)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(404)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var res models.GroupDocument
	var jobPayload webhook.JobsEvent

	bb, _ := bson.Marshal(findRes)
	bson.Unmarshal(bb, &res)

	body, _ := io.ReadAll(r.Body)
	var pay models.ObjectKind
	err = json.Unmarshal(body, &pay)
	if err != nil {
		log.Panicln(err)
	}

	switch pay.ObjectKind {
	case "build":
		err = json.Unmarshal(body, &jobPayload)
		if err != nil {
			panic(err)
		}

		configs.SetActualJob(jobPayload.Commit.ID)
		if jobPayload.Commit.ID == configs.GetCurrentJob() {
			EditMessage(res.ChatId, pay, body)
		}
	case "push":
		configs.SetCurrentJob(configs.GetActualJob())
		currentMesID++
		SendMessage(res.ChatId, pay, body)
	case "pipeline":
		break
	default:
		SendMessage(res.ChatId, pay, body)
	}
}
