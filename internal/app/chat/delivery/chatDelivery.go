package delivery

import (
	"github.com/Andronovdima/golang-chat/internal/app/admin/general"
	"github.com/Andronovdima/golang-chat/internal/app/chat"
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

	m.HandleFunc("/", handler.HandleChat).Methods(http.MethodPut, http.MethodOptions)
	m.Handle("/" , )
}


func (c *ChatHandler) HandleChat(w http.ResponseWriter, r *http.Request) {

	general.Respond(w, r , http.StatusOK, "Hello from world")
}
