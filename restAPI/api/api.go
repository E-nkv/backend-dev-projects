package api

import (
	"log"

	"github.com/E-nkv/backend-dev-projects/restAPI/service"
	"github.com/go-chi/chi/v5"
)

type App struct {
	//keep here all required dependencies, instead of having them globally.
	Service service.Service
	Log     *log.Logger
}

func (app *App) Mount(r *chi.Mux) {
	//all the handlers here
	r.Get("/", app.HandleHome)
	r.Route("/users", func(r chi.Router) {
		r.Get("/", app.HandleGetUsers)
		r.Get("/{userID}", app.HandleGetUser)
		r.Post("/", app.HandleCreateUser)
		r.Delete("/{userID}", app.HandleDeleteUser)
	})

}
