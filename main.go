package main

import (
	"fmt"

	"github.com/woaitsAryan/MyMandap-task/helpers"
	"github.com/woaitsAryan/MyMandap-task/models"
	"github.com/woaitsAryan/MyMandap-task/utils"
)

func init(){
	models.SeedUsers(10)
}

func main() {
    for {
        fmt.Println("1. Send message between 2 users")
        fmt.Println("2. Broadcast message to all users")
        fmt.Println("3. View message log")
        fmt.Println("4. Exit")
        choice := utils.GetInputString("Enter choice: ")
        switch choice {
        case "1":
            senderID := utils.GetInputInt("Enter sender ID: ")
			receiverID := utils.GetInputInt("Enter receiver ID: ")
			content := utils.GetInputString("Enter message: ")
			helpers.SendMessage(senderID, receiverID, content)
        case "2":
			senderID := utils.GetInputInt("Enter sender ID: ")
			content := utils.GetInputString("Enter message: ")
			helpers.BroadcastMessage(senderID, content)
        case "3":
			userID := utils.GetInputInt("Enter sender ID: ")
            helpers.ViewLog(userID)
        case "4":
            for _, user := range models.Users {
                helpers.ViewLog(user.ID)
            }
            return
        default:
            fmt.Println("Invalid choice")
        }
    }
}