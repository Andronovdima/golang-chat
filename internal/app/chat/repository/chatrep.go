package repository

import (
	"database/sql"
	"github.com/Andronovdima/golang-chat/internal/app/chat"
	"github.com/Andronovdima/golang-chat/internal/model"
)

type ChatRepository struct {
	db *sql.DB
}

func NewChatRepository(db *sql.DB) chat.Repository {
	return &ChatRepository{db}
}

func (r *ChatRepository) Create(chat *model.Chat) error {
	return r.db.QueryRow(
		"INSERT INTO chats (user_id, name) " +
			"VALUES ($1, $2) RETURNING id",
		chat.UserID,
		chat.Name,
	).Scan(&chat.ID)
}

func (r *ChatRepository) Find(id int64) (*model.Chat, error) {
	c := &model.Chat{}
	if err := r.db.QueryRow(
		"SELECT id, user_id, support_id, name FROM chats WHERE id = $1",
		id,
	).Scan(
		&c.ID,
		&c.UserID,
		&c.SupportID,
		&c.Name,
	); err != nil {
		return nil, err
	}
	return c, nil
}

func (r *ChatRepository) FindByUser(id int64) (*model.Chat, error) {
	c := &model.Chat{}
	if err := r.db.QueryRow(
		"SELECT id, user_id, support_id, name FROM chats WHERE user_id = $1",
		id,
	).Scan(
		&c.ID,
		&c.UserID,
		&c.SupportID,
		&c.Name,
	); err != nil {
		return nil, err
	}
	return c, nil
}

func (r *ChatRepository) List() ([]*model.Chat, error) {
	var chats []*model.Chat
	rows, err := r.db.Query(
		"SELECT id, user_id, support_id, name FROM chats ORDER BY id DESC")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		c := &model.Chat{}
		err := rows.Scan(&c.ID, &c.UserID, &c.SupportID, &c.Name)
		if err != nil {
			return nil, err
		}
		chats = append(chats, c)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	return chats, nil
}

