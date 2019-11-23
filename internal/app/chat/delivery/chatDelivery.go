package delivery

import (
	"Golang/FworkChat/golang-chat/internal/app/chat"
	"github.com/gorilla/mux"
	"net/http"
)

type ChatHandler struct {
	ChatUsecase		chat.Usecase
}

func NewChatHandler(m *mux.Router, uc chat.Usecase) {
	handler := &ChatHandler{
		ChatUsecase:	uc,
	}

	m.HandleFunc("/chat", handler.HandleChat).Methods(http.MethodPut, http.MethodOptions)
}


func (c *ChatHandler) HandleChat(w http.ResponseWriter, r *http.Request) {
// TODO
}
