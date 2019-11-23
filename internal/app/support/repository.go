package support

import "github.com/Andronovdima/golang-chat/internal/model"

type Repository interface {
	Create(support *model.Support) error
	Find(id int64) (*model.Support, error)
}
