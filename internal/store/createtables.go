package store

import "database/sql"
func CreateTables(db *sql.DB) error {
	supportsQuery := `CREATE TABLE IF NOT EXISTS supports (
		id bigserial not null primary key,
		user_id integer
	);`

	if _, err := db.Exec(supportsQuery); err != nil {
		return err
	}

	adminsQuery := `CREATE TABLE IF NOT EXISTS admins (
		id bigserial not null primary key,
		user_id integer
	);`

	if _, err := db.Exec(adminsQuery); err != nil {
		return err
	}

	if _, err := db.Exec(adminsQuery); err != nil {
		return err
	}

	chatsQuery := `CREATE TABLE IF NOT EXISTS chats (
		id bigserial not null primary key,
		name varchar,
		user_id integer not null,
		support_id integer
	);`

	if _, err := db.Exec(chatsQuery); err != nil {
		return err
	}

	messagesQuery := `CREATE TABLE IF NOT EXISTS messages (
		id bigserial not null primary key,
		message varchar,
		sender_id integer not null,
		receiver_id integer,
		date timestamp
	);`

	if _, err := db.Exec(messagesQuery); err != nil {
		return err
	}
	return nil
}