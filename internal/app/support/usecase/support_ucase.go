package supportusecase

import (
	"github.com/Andronovdima/golang-chat/internal/app/support"
	"github.com/Andronovdima/golang-chat/internal/model"
	"github.com/pkg/errors"
)

type SupportUsecase struct {
	supportRep support.Repository
}

func NewSupportUsecase(r support.Repository) support.Usecase {
	return &SupportUsecase{
		supportRep: r,
	}
}

func (u *SupportUsecase) Find(id int64) (*model.Support, error) {
	s, err := u.supportRep.Find(id)
	if err != nil {
		return nil, errors.Wrapf(err, "supportRep.Find()")
	}
	return s, nil
}