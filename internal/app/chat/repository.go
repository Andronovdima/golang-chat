package chat

import "github.com/Andronovdima/golang-chat/internal/model"

type Repository interface {
	Create(chat *model.Chat) error
	Find(id int64) (*model.Chat, error)
	FindByUser(id int64) (*model.Chat, error)
	List() ([]*model.Chat, error)
}
