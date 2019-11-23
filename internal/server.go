package internal

import (
	"database/sql"
	"encoding/json"
	"github.com/Andronovdima/golang-chat/internal/app/chat"
	repository2 "github.com/Andronovdima/golang-chat/internal/app/chat/repository"
	chatusecase "github.com/Andronovdima/golang-chat/internal/app/chat/usecase"
	"github.com/Andronovdima/golang-chat/internal/app/message"
	"github.com/Andronovdima/golang-chat/internal/app/message/repository"
	"github.com/Andronovdima/golang-chat/internal/app/message/usecase"
	"github.com/Andronovdima/golang-chat/internal/model"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// Chat server.
type Server struct {
	pattern   string
	//messages  []*Message
	clients   map[int]*Client
	addCh     chan *Client
	delCh     chan *Client
	sendAllCh chan *model.Message
	doneCh    chan bool
	errCh     chan error
	MesUcase  message.Usecase
	ChatUcase chat.Usecase
}

// Create new app server.
func NewServer(pattern string, db *sql.DB) *Server {
	//messages := []*Message{}
	clients := make(map[int]*Client)
	addCh := make(chan *Client)
	delCh := make(chan *Client)
	sendAllCh := make(chan *model.Message)
	doneCh := make(chan bool)
	errCh := make(chan error)

	return &Server{
		pattern,
		clients,
		addCh,
		delCh,
		sendAllCh,
		doneCh,
		errCh,
		usecase.NewMessageUsecase(repository.NewMessageRepository(db)),
		chatusecase.NewChatUsecase(repository2.NewChatRepository(db)),
	}
}

func (s *Server) Add(c *Client) {
	s.addCh <- c
}

func (s *Server) Del(c *Client) {
	s.delCh <- c
}

func (s *Server) SendAll(msg *model.Message) {
	s.sendAllCh <- msg
}

func (s *Server) Done() {
	s.doneCh <- true
}

func (s *Server) Err(err error) {
	s.errCh <- err
}

func (s *Server) sendPastMessages(c *Client) {
	messages, _ := s.MesUcase.ListByUser(c.userId)
	for _, msg := range messages {
		c.Write(msg)
	}
}

func (s *Server) sendAll(msg *model.Message) {
	log.Println("send All:", msg)
	for _, c := range s.clients {
		c.Write(msg)
	}
}

// Listen and serve.
// It serves client connection and broadcast request.
func (s *Server) Listen() {

	log.Println("Listening server...")

	// websocket handler
	onConnected := func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				s.errCh <- err
			}
		}()

		client := NewClient(ws, s)
		s.Add(client)
		client.Listen()
	}
	http.Handle(s.pattern, websocket.Handler(onConnected))
	log.Println("Created handler")

	for {
		select {

		// Add new a client
		case c := <-s.addCh:
			log.Println("Added new client")
			s.clients[c.id] = c
			log.Println("Now", len(s.clients), "clients connected.")
			s.sendPastMessages(c)

		// del a client
		case c := <-s.delCh:
			log.Println("Delete client")
			delete(s.clients, c.id)

		// broadcast message for all clients
		case msg := <-s.sendAllCh:
			log.Println("Send all:", msg)
			_ = s.MesUcase.Create(msg)
			//s.messages = append(s.messages, msg)
			s.sendAll(msg)

		case err := <-s.errCh:
			log.Println("Error:", err.Error())

		case <-s.doneCh:
			return
		}
	}
}

func (s *Server) MainHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	str := "hello from server"
	if &str != nil {
		_ = json.NewEncoder(w).Encode(&str)
	}
}