package exercise

import (
	"encoding/json"
	"net/http"

	"github.com/go-boyars/gym/internal/models"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserAPI struct {
	Login    string `json:"login"`
	Password string `json:"pass"`
}

func (a *App) registerHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var userApi RegisterUserAPI
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

	_, err = response.Write([]byte(`{"id": "` + id + `"}`))
	if err != nil {
		// and what? log it
		_ = err
	}
}

type LoginAPI struct {
	Login    string `json:"login"`
	Password string `json:"pass"`
}

func (a *App) loginHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var loginApi LoginAPI
	err := json.NewDecoder(request.Body).Decode(&loginApi)
	if err != nil {
		SetInternalError(response, err)
		return
	}

	actualHash, err := a.storage.GetPwhash(request.Context(), loginApi.Login)
	if err != nil {
		SetInternalError(response, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(actualHash), []byte(loginApi.Password))
	if err != nil {
		response.WriteHeader(http.StatusForbidden)
		_, err := response.Write([]byte(`{"message": "invalid password"}`))
		if err != nil {
			// and what? log it
			_ = err
		}
		return
	}
}
