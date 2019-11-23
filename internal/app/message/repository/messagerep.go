package repository

import (
	"database/sql"
	"github.com/Andronovdima/golang-chat/internal/app/message"
	"github.com/Andronovdima/golang-chat/internal/model"
)

type MessageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) message.Repository {
	return &MessageRepository{db}
}

func (r *MessageRepository) Create(message *model.Message) error {
	return r.db.QueryRow(
		"INSERT INTO messages (chat_id, sender_id, receiver_id, message, date) " +
			"VALUES ($1, $2, $3, $4) RETURNING id",
		message.ChatID,
		message.SenderID,
		message.ReceiverID,
		message.Message,
		message.Date,
	).Scan(&message.ID)
}

func (r *MessageRepository) ListByUser(id int64) ([]*model.Message, error) {
	var messages []*model.Message
	rows, err := r.db.Query(
		"SELECT id, chat_id, sender_id, receiver_id, message, date FROM messages " +
			"WHERE sender_id = $1 OR receiver_id = $1 ORDER BY date DESC", id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		m := &model.Message{}
		err := rows.Scan(&m.ID, &m.ChatID, &m.SenderID, &m.ReceiverID, &m.Message, &m.Date)
		if err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *MessageRepository) ListBySupport(supportId int64, chatId int64) ([]*model.Message, error) {
	var messages []*model.Message
	rows, err := r.db.Query(
		"SELECT id, chat_id, sender_id, receiver_id, message, date FROM messages " +
			"WHERE (sender_id = $1 OR receiver_id = $1) AND chat_id = $2 ORDER BY date DESC", supportId, chatId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		m := &model.Message{}
		err := rows.Scan(&m.ID, &m.ChatID, &m.SenderID, &m.ReceiverID, &m.Message, &m.Date)
		if err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	return messages, nil
}