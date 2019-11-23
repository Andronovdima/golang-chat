package model

import (
	"strconv"
	"time"
)

type Message struct {
	ID 			int64
	ChatID 		int64
	SenderID 	int64
	ReceiverID 	int64
	Message 	string
	Date		time.Time
}

func (self *Message) String() string {
	return strconv.Itoa(int(self.SenderID)) + " says " + self.Message
}