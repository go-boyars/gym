package exercise

import (
	"encoding/json"
	"net/http"

	"github.com/go-boyars/gym/internal/models"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserAPI struct {
	Login    string `json:"login"`
	Password string `json:"pass"`
}

func (a *App) registerHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var userApi CreateUserAPI
	err := json.NewDecoder(request.Body).Decode(&userApi)
	if err != nil {
		SetInternalError(response, err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userApi.Password), 10)
	if err != nil {
		SetInternalError(response, err)
		return
	}

	id := uuid.NewV4().String()
	user := models.User{
		ID:    id,
		Login: userApi.Login,
	}

	err = a.storage.CreateUser(request.Context(), user, string(hash))
	if err != nil {
		SetInternalError(response, err)
		return
	}
}
