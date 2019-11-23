package internal

import (
	"database/sql"
	"flag"
	"github.com/Andronovdima/golang-chat/internal/store"
	_ "github.com/lib/pq"
	"log"
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

func Start() {
	flag.Parse()

	config := NewConfig()
	db, _ := newDB(config.DatabaseURL)

	log.SetFlags(log.Lshortfile)

	// websocket server
	server := NewServer("/entry", db)
	go server.Listen()

	// static files
	http.Handle("/", http.FileServer(http.Dir("webroot")))
	//http.Handle("/hello", )

	log.Fatal(http.ListenAndServe(":8080", nil))
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
