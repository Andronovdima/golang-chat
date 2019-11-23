package adminrepository

import (
	"database/sql"
	"github.com/Andronovdima/golang-chat/internal/model"
)
import "github.com/Andronovdima/golang-chat/internal/app/admin"

type AdminRepository struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) admin.Repository {
	return &AdminRepository{db}
}

func (r *AdminRepository) Create(admin *model.Admin) error {
	return r.db.QueryRow(
		"INSERT INTO admins (id, user_id) " +
			"VALUES ($1, $2) RETURNING id",
		admin.ID,
		admin.UserID,
	).Scan(&admin.ID)
}

func (r *AdminRepository) Find(id int64) (*model.Admin, error) {
	a := &model.Admin{}
	if err := r.db.QueryRow(
		"SELECT id, user_id FROM admins WHERE id = $1",
		id,
	).Scan(
		&a.ID,
		&a.UserID,
	); err != nil {
		return nil, err
	}
	return a, nil
}
