package model

type Input struct {
	Action string
	Payload Payload
}


type Payload struct{
	AccountID int64
	Message string
}
