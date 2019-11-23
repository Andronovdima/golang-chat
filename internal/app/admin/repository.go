package admin

import "github.com/Andronovdima/golang-chat/internal/model"

type Repository interface {
	Create(admin *model.Admin) error
	Find(id int64) (*model.Admin, error)
}
