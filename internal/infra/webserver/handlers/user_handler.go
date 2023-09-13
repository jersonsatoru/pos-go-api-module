package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jersonsatoru/pos-go-api-module/internal/entities"
	"github.com/jersonsatoru/pos-go-api-module/internal/infra/database"
	"github.com/jersonsatoru/pos-go-api-module/internal/infra/webserver/dtos"
)

type UserHandler struct {
	Repository database.UserRepository
}

func NewUserHandler(repo database.UserRepository) *UserHandler {
	return &UserHandler{
		Repository: repo,
	}
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input dtos.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entities.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Repository.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
