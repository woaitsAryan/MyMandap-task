package helpers

import (
	"time"

	"github.com/woaitsAryan/MyMandap-task/models"
)

func SendMessage(senderID int, receiverID int, content string) {
	message := models.Message{Sender: senderID, Receiver: receiverID, Content: content, Timestamp: time.Now()}
	models.Users[senderID].Outbox = append(models.Users[senderID].Outbox, message)
	models.Users[receiverID].Inbox = append(models.Users[receiverID].Inbox, message)
}

func BroadcastMessage(senderID int, content string) {
	for id := range models.Users {
		if id != senderID {
			SendMessage(senderID, id, content)
		}
	}
}