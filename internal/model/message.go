package model

import "time"

type Message struct {
	ID 			int64
	ChatID 		int64
	SenderID 	int64
	ReceiverID 	int64
	Message 	string
	Date		time.Time
}
