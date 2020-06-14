package exercise

import (
	"encoding/json"
	"net/http"
)

type Exercise struct {
	Id      int
	Name    string
	Muscule string
}

func (a *App) getExercises(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

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
