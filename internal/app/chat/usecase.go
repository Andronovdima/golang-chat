package chat

import "github.com/Andronovdima/golang-chat/internal/model"

type Usecase interface {
	Create(userId int64) (*model.Chat, error)
	Find(id int64) (*model.Chat, error)
	FindByUser(id int64) (*model.Chat, error)
	List() ([]*model.Chat, error)
}