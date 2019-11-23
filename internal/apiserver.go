package internal

import (
	"database/sql"
	"github.com/Andronovdima/golang-chat/internal/store"
	_ "github.com/lib/pq"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	code int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func Start() error {
	config := NewConfig()
	db, _ := newDB(config.DatabaseURL)
	server := NewServer("/entry", db)
	go server.Listen()
	return http.ListenAndServe(":8080", nil)
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(20)
	if err := store.CreateTables(db); err != nil {
		return nil, err
	}
	return db, nil
}
