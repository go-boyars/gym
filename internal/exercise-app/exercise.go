package exercise

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/context"
)

type Exercise struct {
	Id      int
	Name    string
	Muscule string
}

func (a *App) getExercises(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	claims := context.Get(request, "jwt").(CustomJWTClaim)
	fmt.Println("user id: ", claims.ID)

	// TODO context
	exercises, err := a.storage.GetExercises()
	if err != nil {
		SetInternalError(response, err)
		return
	}

	err = json.NewEncoder(response).Encode(exercises)
	if err != nil {
		SetInternalError(response, err)
		return
	}
}
