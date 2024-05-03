package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/woaitsAryan/MyMandap-task/helpers"
	"github.com/woaitsAryan/MyMandap-task/models"
	"github.com/woaitsAryan/MyMandap-task/utils"
)

func init(){
	models.SeedUsers(100)
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Println("1. Send message between 2 users")
        fmt.Println("2. Broadcast message to all users")
        fmt.Println("3. View message log")
        fmt.Println("4. Exit")
        fmt.Print("Enter choice: ")
        choice, _ := reader.ReadString('\n')
        choice = strings.TrimSpace(choice)
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