package chatusecase

import (
	"github.com/Andronovdima/golang-chat/internal/app/chat"
	"github.com/Andronovdima/golang-chat/internal/model"
	"github.com/pkg/errors"
)

type ChatUsecase struct {
	rep chat.Repository
}

func NewChatUsecase(r chat.Repository) chat.Usecase {
	return &ChatUsecase{
		rep: r,
	}
}


func (u *ChatUsecase) Create(userId int64) (*model.Chat, error) {
	c := &model.Chat{
		UserID: userId,
	}
	if err := u.rep.Create(c); err != nil {
		return nil, errors.Wrapf(err, "chatRep.Create()")
	}
	return c, nil
}

func (u *ChatUsecase) Find(id int64) (*model.Chat, error) {
	c, err := u.rep.Find(id)
	if err != nil {
		return nil, errors.Wrapf(err, "chatRep.Find()")
	}
	return c, nil
}

func (u *ChatUsecase) FindByUser(userId int64) (*model.Chat, error) {
	c, err := u.rep.FindByUser(userId)
	if err != nil {
		return nil, errors.Wrapf(err, "chatRep.Find()")
	}
	return c, nil
}

func (u *ChatUsecase) List() ([]*model.Chat, error) {
	chats, err := u.rep.List()
	if err != nil {
		return nil, errors.Wrap(err, "chatRep.List()")
	}
	return chats, nil
}
