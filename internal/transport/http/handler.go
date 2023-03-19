package http

import (
	"encoding/json"
	"fmt"
	"github.com/Solblnc/Rest-API/internal/coment"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Handler - struct of pointers for services
type Handler struct {
	Router  *mux.Router
	Service *coment.Service
}

// Response - struct for responses of our API
type Response struct {
	Message string
	Error   string
}

// NewHandler Initializing a new Handler
func NewHandler(service *coment.Service) *Handler {
	return &Handler{Service: service}
}

// SetUpRoutes Set up all the routes for application ...
func (h *Handler) SetUpRoutes() {
	fmt.Println("Routes are setted")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=uint8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I am alive!"}); err != nil {
			panic(err)
		}
	})
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=uint8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendError(w, "Error in parsing id to uint", err)
		return
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		sendError(w, "Error in retrieving comment by id", err)
		return
	}

	if err = json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}

}

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=uint8")
	w.WriteHeader(http.StatusOK)

	comments, err := h.Service.GetAllComments()
	if err != nil {
		sendError(w, "Error of retrieve all comments", err)
		return
	}

	if err = json.NewEncoder(w).Encode(comments); err != nil {
		panic(err)
	}

}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=uint8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendError(w, "Error in parsing id to uint", err)
		return
	}

	err = h.Service.DeleteComment(uint(i))
	if err != nil {
		sendError(w, "Error in deleting a comment", err)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Comment deleted"}); err != nil {
		panic(err)
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=uint8")
	w.WriteHeader(http.StatusOK)

	var comment coment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendError(w, "Failed to decode a json body", err)
		return
	}

	comment, err := h.Service.PostComment(comment)
	if err != nil {
		sendError(w, "Failed to post a new comment", err)
		return
	}

	if err = json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=uint8")
	w.WriteHeader(http.StatusOK)

	var comment coment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendError(w, "Failed to decode a json body", err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendError(w, "Error in parsing id to uint", err)
		return
	}

	comment, err = h.Service.UpdateComment(uint(i), comment)
	if err != nil {
		sendError(w, "Failed to update a comment", err)
		return
	}

	if err = json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

func sendError(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
