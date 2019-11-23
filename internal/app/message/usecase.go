package message

import "github.com/Andronovdima/golang-chat/internal/model"

type Usecase interface {
	Create(message *model.Message) error
	ListByUser(id int64) ([]*model.Message, error)
	ListBySupport(supportId int64, chatId int64) ([]*model.Message, error)
}
