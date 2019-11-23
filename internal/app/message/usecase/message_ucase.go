package usecase

import (
	"github.com/Andronovdima/golang-chat/internal/app/message"
	"github.com/Andronovdima/golang-chat/internal/model"
	"github.com/pkg/errors"
	"time"
)

type MessageUsecase struct {
	rep message.Repository
}

func NewMessageUsecase(r message.Repository) message.Usecase {
	return &MessageUsecase{
		rep: r,
	}
}

func (u *MessageUsecase) Create(mes *model.Message) error {
	mes.Date = time.Now()
	if err := u.rep.Create(mes); err != nil {
		return errors.Wrap(err, "messageRep.Create()")
	}
	return nil
}

func (u *MessageUsecase) ListByUser(userId int64) ([]*model.Message, error) {
	messages, err := u.rep.ListByUser(userId)
	if err != nil {
		return nil, errors.Wrap(err, "messageRep.ListByUser()")
	}
	return messages, nil
}

func (u *MessageUsecase) ListBySupport(supportId int64, chatId int64) ([]*model.Message, error) {
	messages, err := u.rep.ListBySupport(supportId, chatId)
	if err != nil {
		return nil, errors.Wrap(err, "messageRep.ListBySupport()")
	}
	return messages, nil
}

