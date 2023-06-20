package configs

import (
	"gitbot/models"
	"log"
	"os"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	lock        = &sync.Mutex{}
	env         *models.Config
	checkStatus primitive.ObjectID
	pushMessage tgbotapi.Message
	actualJob   int
	currentJob  int
)

func GetConfig() *models.Config {
	if env == nil {
		lock.Lock()
		defer lock.Unlock()

		if env == nil {
			if err := godotenv.Load(); err != nil {
				log.Printf(".env file not found.\n")
			}

			env = &models.Config{
				Port:          os.Getenv("PORT"),
				HostURL:       os.Getenv("HOST_URL"),
				PathURL:       os.Getenv("URL_PATH"),
				BotToken:      os.Getenv("TELEGRAM_BOT_TOKEN"),
				MongoURI:      os.Getenv("MONGO_URI"),
				MongoDatabase: os.Getenv("MONGO_DATABASE"),
			}

			log.Println("Configurations loaded successfully.")
		}
	}

	return env
}

func GetCheckStatus() primitive.ObjectID {
	return checkStatus
}
func SetCheckStatus(objectId primitive.ObjectID) {
	checkStatus = objectId
}

func GetPushMessage() tgbotapi.Message {
	return pushMessage
}
func SetPushMessage(msg tgbotapi.Message, content string) {
	msg.Text = content
	pushMessage = msg
}

func GetActualJob() int {
	return actualJob
}
func SetActualJob(value int) {
	actualJob = value
}

func GetCurrentJob() int {
	return currentJob
}
func SetCurrentJob(value int) {
	currentJob = value
}
