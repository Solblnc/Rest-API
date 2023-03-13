package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Handler - struct of pointers for services
type Handler struct {
	Router *mux.Router
}

// NewHandler Initializing a new Handler
func NewHandler() *Handler {
	return &Handler{}
}

// SetUpRoutes Set up all the routes for application ...
func (h *Handler) SetUpRoutes() {
	fmt.Println("Routes are setted")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive!")
	})
}
