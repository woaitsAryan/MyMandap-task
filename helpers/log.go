package helpers

import (
	"fmt"

	"github.com/woaitsAryan/MyMandap-task/models"
)

func ViewLog(id int) {

	if  isInboxEmpty(models.Users[id]) && isOutboxEmpty(models.Users[id]) {
		return
	}

    for _, message := range models.Users[id].Inbox {
        fmt.Printf("User %d received a message from User %d at %s: %s\n", id, message.Sender, message.Timestamp.Local().Format("2006-01-02 15:04:05"), message.Content)
    }
	for _, message := range models.Users[id].Outbox {
        fmt.Printf("User %d sent a message to User %d at %s: %s\n", id, message.Receiver, message.Timestamp.Local().Format("2006-01-02 15:04:05"), message.Content)
    }
}


func isInboxEmpty(user *models.User) bool {
	return len(user.Inbox) == 0
}

func isOutboxEmpty(user *models.User) bool {
	return len(user.Outbox) == 0
}