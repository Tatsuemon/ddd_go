package router

import (
	"github.com/Tatsuemon/ddd_go/infrastructure/web/handler"

	"github.com/gorilla/mux"
)

func NewMuxRouter(h handler.AppHandler) *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/users", h.GetUsers).Methods("GET")
	mux.HandleFunc("/users/id/{id}", h.GetUser).Methods("GET")
	mux.HandleFunc("/users", h.CreateUser).Methods("POST")
	mux.HandleFunc("/users/id/{id}", h.UpdateUser).Methods("PUT")
	mux.HandleFunc("/users/id/{id}", h.DeleteUser).Methods("DELETE")
	return mux
}
