package supportrepository

import (
	"database/sql"
	"github.com/Andronovdima/golang-chat/internal/app/support"
	"github.com/Andronovdima/golang-chat/internal/model"
)

type SupportRepository struct {
	db *sql.DB
}

func NewSupportRepository(db *sql.DB) support.Repository {
	return &SupportRepository{db}
}

func (r *SupportRepository) Create(support *model.Support) error {
	return r.db.QueryRow(
		"INSERT INTO supports (user_id) " +
			"VALUES ($1) RETURNING id",
		support.UserID,
	).Scan(&support.ID)
}

func (r *SupportRepository) Find(id int64) (*model.Support, error) {
	s := &model.Support{}
	if err := r.db.QueryRow(
		"SELECT id, user_id FROM supports WHERE id = $1",
		id,
	).Scan(
		&s.ID,
		&s.UserID,
	); err != nil {
		return nil, err
	}
	return s, nil
}
