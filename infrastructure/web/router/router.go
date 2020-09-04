package router

import (
	"ddd_go/infrastructure/web/handler"

	"github.com/gorilla/mux"
)

func NewMuxRouter() *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/users", handler.UserHandler.GetUsers).Methods("GET")
	mux.HandleFunc("/users/{id}", handler.UserHandler.GetUser).Methods("GET")
	mux.HandleFunc("/users", handler.UserHandler.CreateUser).Methods("POST")
	mux.HandleFunc("/users/{id}", handler.UserHandler.UpdateUser).Methods("PUT")
	mux.HandleFunc("/users/{id}", handler.UserHandler.DeleteUsers).Methods("DELETE")
	return mux
}
