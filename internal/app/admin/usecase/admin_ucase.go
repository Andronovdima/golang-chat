package adminucase

import (
	"github.com/Andronovdima/golang-chat/internal/app/admin"
	"github.com/Andronovdima/golang-chat/internal/model"
	"github.com/pkg/errors"
)

type AdminUsecase struct {
	adminRep admin.Repository
}

func NewAdminUsecase(c admin.Repository) admin.Usecase {
	return &AdminUsecase{
		adminRep: c,
	}
}

func (u *AdminUsecase) Find(id int64) (*model.Admin, error) {
	a, err := u.adminRep.Find(id)
	if err != nil {
		return nil, errors.Wrapf(err, "adminRep.Find()")
	}
	return a, nil
}