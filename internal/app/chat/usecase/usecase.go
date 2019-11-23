package usecase

import "github.com/Andronovdima/golang-chat/internal/app/chat"

type chatUsecase struct {
	a int
}

func NewChatUsecase() chat.Usecase {
	return &chatUsecase{
		a: 5,
	}
}