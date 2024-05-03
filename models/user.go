package models

import (
	"time"
)

type Message struct {
	Sender    int
	Receiver  int
	Content   string
	Timestamp time.Time
}

type User struct {
	ID     int
	Inbox  []Message
	Outbox []Message
}

var Users = make(map[int]*User)

func SeedUsers(number int) {
	for i := 1; i <= number; i++ {
		Users[i] = &User{ID: i}
	}
}