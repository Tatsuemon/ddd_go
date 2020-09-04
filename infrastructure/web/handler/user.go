package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Tatsuemon/ddd_go/domain/model"
	"github.com/Tatsuemon/ddd_go/usecase"

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

func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return &userHandler{u}
}

func (h *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	u, err := h.UserUseCase.GetUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]interface{}{"users": u}
	respondWithJson(w, http.StatusOK, response)
	return
}

func (h *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	u, err := h.UserUseCase.GetUser(id)
	if err != nil {
		// TODO: not foundもある
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]interface{}{"user": u}
	respondWithJson(w, http.StatusOK, response)
	return
}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	type payload struct {
		Name  string `json: "name"`
		Email string `json: "email"`
	}
	var p payload
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	u, err := model.NewUser(
		p.Name,
		p.Email,
	)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	u, err = h.UserUseCase.CreateUser(u)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response := map[string]interface{}{"user": u}
	respondWithJson(w, http.StatusCreated, response)
	return
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var u *model.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	u, err := h.UserUseCase.UpdateUser(u, vars["id"])

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response := map[string]interface{}{"user": u}
	respondWithJson(w, http.StatusOK, response)
	return
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.UserUseCase.DeleteUser(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
	return
}

// TODO(Tatsuemon): 別で切り出す
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJson(w, code, map[string]string{"error": message})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	// TODO(Tatsuemon): 変える
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
