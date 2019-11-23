package model

type Input struct {
	Action string
	Payload Payload
}


type Payload struct{
	AccountId int64
	Message Message
}
