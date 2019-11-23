package apiserver

import (
	"database/sql"
	"github.com/Andronovdima/golang-chat/internal/app/chat"
	"github.com/Andronovdima/golang-chat/internal/model"
	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	chUsecase "github.com/Andronovdima/golang-chat/internal/app/chat/usecase"
	chDelivery "github.com/Andronovdima/golang-chat/internal/app/chat/delivery"
)

type Server struct {
	Mux          *mux.Router
	Config       *Config
	clients   map[int]*chat.Client
	addCh     chan *chat.Client
	delCh     chan *chat.Client
	sendAllCh chan *model.Message
	doneCh    chan bool
	errCh     chan error
}

func NewServer(config *Config , ) (*Server, error) {
	clients := make(map[int]*chat.Client)
	addCh := make(chan *chat.Client)
	delCh := make(chan *chat.Client)
	sendAllCh := make(chan *model.Message)
	doneCh := make(chan bool)
	errCh := make(chan error)

	s := &Server{
		mux.NewRouter(),
		config,
		clients,
		addCh,
		delCh,
		sendAllCh,
		doneCh,
		errCh,
	}
	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Mux.ServeHTTP(w, r)
}



func (s *Server) ConfigureServer(db *sql.DB) {
	uc := chUsecase.NewChatUsecase()
	chDelivery.NewChatHandler(s.Mux, uc)
//	TODO confugure server
}

func (s *Server) Add(c *chat.Client) {
	s.addCh <- c
}

func (s *Server) Del(c *chat.Client) {
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

func (s *Server) sendPastMessages(c *chat.Client) {
	for _, msg := range s.messages {
		c.Write(msg)
	}
}

func (s *Server) sendAll(msg *model.Message) {
	for _, c := range s.clients {
		c.Write(msg)
	}
}


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

		client := chat.NewClient(ws, s)
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
			s.messages = append(s.messages, msg)
			s.sendAll(msg)

		case err := <-s.errCh:
			log.Println("Error:", err.Error())

		case <-s.doneCh:
			return
		}
	}
}
