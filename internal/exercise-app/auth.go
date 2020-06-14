package exercise

import (
	"encoding/json"
	"net/http"

	"github.com/go-boyars/gym/internal/models"
)

func (a *App) registerHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		SetInternalError(response, err)
		return
	}
	hash := "soeifj" // TODO

	err = a.storage.CreateUser(request.Context(), user, hash)
	if err != nil {
		SetInternalError(response, err)
		return
	}
}
