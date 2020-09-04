package handler

import (
	"encoding/json"
	"github.com/Tatsuemon/ddd_go/domain/model"
	"github.com/Tatsuemon/ddd_go/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	usecase.UserUseCase
}

func NewUserHandler(u usecase.UserHandler) UserHandler {
	return &userHandler{u}
}

func (h *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := h.UserUseCase.GetUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	return respondWithJson(w, http.StatusOK, users)
}

func (h *userHandler) GetUser(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]
	user, err := h.UserUseCase.GetUser(id)
	if err != nil {
		// TODO: not foundもある
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	return respondWithJson(w, http.StatusOK, user)
}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	var u model.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if u, err := h.UserUseCase.CreateUser(u); err != nil {
		respondWithError(w, htp.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, u)
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	var u model.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()
	vars := mux.Vars(r)
	id := vars["id"]
	u.ID = id

	if u, err := h.UserUseCase.UpdateUser(u); err != nil {
		respondWithError(w, htp.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, u)
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]
	u.ID = id

	if err := h.UserUseCase.DeleteUser(u); err != nil {
		respondWithError(w, htp.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// TODO(Tatsuemon): 別で切り出す
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJson(w, code, map[string]string{"error": message})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
