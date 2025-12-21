package user

import "net/http"

type Handler struct {
}

func NewHanlder() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /login", h.HandleLogin)
	router.HandleFunc("POST /register", h.HandleRegister)
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login"))
}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("register"))
}
