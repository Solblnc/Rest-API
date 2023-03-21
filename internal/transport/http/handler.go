package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Solblnc/Rest-API/internal/coment"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
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

// LoggingMiddleware - adds middleware to endpoints
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"Method": r.Method,
			"Path":   r.URL.Path,
		}).Info("Handled request")
		next.ServeHTTP(w, r)
	})
}

// BasicAuth - provide a basic auth
func BasicAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("Basic auth endpoint hit")
		user, pass, ok := r.BasicAuth()
		if user == "admin" && pass == "password" && ok {
			original(w, r)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			sendError(w, "not authorized", errors.New("Not authorized"))
		}

	}
}

func validateToken(accessToken string) bool {
	var mySignKey = []byte("mission")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there has been an error")
		}
		return mySignKey, nil
	})

	if err != nil {
		return false
	}

	return token.Valid
}

// JWTAuth - a handy middleware function that will provide basic auth around specific endpoints
func JWTAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("jwt auth endpoint hit")
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			sendError(w, "not authorized", errors.New("not authorized"))
		}

		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			sendError(w, "not authorized", errors.New("not authorized"))
		}

		if validateToken(authHeaderParts[1]) {
			original(w, r)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			sendError(w, "not authorized", errors.New("not authorized"))
		}
	}
}

// SetUpRoutes Set up all the routes for application ...
func (h *Handler) SetUpRoutes() {
	fmt.Println("Routes are setted")
	h.Router = mux.NewRouter()
	h.Router.Use(LoggingMiddleware)

	h.Router.HandleFunc("/api/comment", JWTAuth(h.PostComment)).Methods("POST")
	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", BasicAuth(h.DeleteComment)).Methods("DELETE")
	h.Router.HandleFunc("/api/comment/{id}", BasicAuth(h.UpdateComment)).Methods("PUT")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=uint8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I am alive!"}); err != nil {
			panic(err)
		}
	})
}

func sendError(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
