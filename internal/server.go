package internal

import (
	"database/sql"
	"github.com/Andronovdima/golang-chat/internal/app/admin"
	adminrepository "github.com/Andronovdima/golang-chat/internal/app/admin/repository"
	adminucase "github.com/Andronovdima/golang-chat/internal/app/admin/usecase"
	"github.com/Andronovdima/golang-chat/internal/app/chat"
	"github.com/Andronovdima/golang-chat/internal/app/chat/repository"
	chatusecase "github.com/Andronovdima/golang-chat/internal/app/chat/usecase"
	"github.com/Andronovdima/golang-chat/internal/app/message"
	repository2 "github.com/Andronovdima/golang-chat/internal/app/message/repository"
	"github.com/Andronovdima/golang-chat/internal/app/message/usecase"
	"github.com/Andronovdima/golang-chat/internal/app/support"
	supportrepository "github.com/Andronovdima/golang-chat/internal/app/support/repository"
	supportusecase "github.com/Andronovdima/golang-chat/internal/app/support/usecase"
	"github.com/Andronovdima/golang-chat/internal/model"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

// Chat server.
type Server struct {
	pattern   string
	clients   map[int]*Client
	addCh     chan *Client
	delCh     chan *Client
	sendAllCh chan *model.Message
	doneCh    chan bool
	errCh     chan error
	AdminUcase	admin.Usecase
	ChatUcase	chat.Usecase
	MessageUcase	message.Usecase
	SupportUcase	support.Usecase
}

// Create new app server.
func NewServer(pattern string, db *sql.DB) *Server {
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
		adminucase.NewAdminUsecase(adminrepository.NewAdminRepository(db)),
		chatusecase.NewChatUsecase(repository.NewChatRepository(db)),
		usecase.NewMessageUsecase(repository2.NewMessageRepository(db)),
		supportusecase.NewSupportUsecase(supportrepository.NewSupportRepository(db)),
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
	messages, err := s.MessageUcase.ListByUser(int64(1))
	if err != nil {
		c.Done()
	}

	for _, msg := range messages {
		c.Write(msg)
	}
}

func (s *Server) sendAll(msg *model.Message) {
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
			msg.SenderID = 1
			err := s.MessageUcase.Create(msg)
			log.Println("Created in db:", msg)
			if err != nil {
				s.errCh <- err
			}
			s.sendAll(msg)

		case err := <-s.errCh:
			log.Println("Error:", err.Error())

		case <-s.doneCh:
			return
		}
	}
}
