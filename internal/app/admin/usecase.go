package admin

import "github.com/Andronovdima/golang-chat/internal/model"

type Usecase interface {
	Find(id int64) (*model.Admin, error)
}
