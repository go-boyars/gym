package exercise

import (
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
	a.router.HandleFunc("/register", a.registerHandler).Methods("POST")
	a.router.HandleFunc("/exercises", a.getExercises).Methods("GET")
}

func SetInternalError(response http.ResponseWriter, handleErr error) {
	response.WriteHeader(http.StatusInternalServerError)
	_, err := response.Write([]byte(`{"message": ` + handleErr.Error() + `}`))
	if err != nil {
		// and what? log it
		_ = err
	}
}
