package exercise

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	storage Storage
	router  *mux.Router
}

func New(storage Storage) (*App, error) {
	r := mux.NewRouter()
	a := &App{
		storage: storage,
		router:  r,
	}

	a.registerHandlers()

	return a, nil
}

func (a *App) Router() *mux.Router {
	return a.router
}

func (a *App) registerHandlers() {
	a.router.HandleFunc("/exercises", a.getExercises).Methods("GET")
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

func SetInternalError(response http.ResponseWriter, handleErr error) {
	response.WriteHeader(http.StatusInternalServerError)
	_, err := response.Write([]byte(`{"message": ` + handleErr.Error() + `}`))
	if err != nil {
		// and what? log it
		_ = err
	}
}
