package main

import (
	"database/sql"
	"log"

	"github.com/go-chi/chi/v5"
)

var port = ":8080"

type App struct {
	//keep here all required dependencies, instead of having them globally.
	DB     *sql.DB
	Logger *log.Logger
}

func (app *App) mount() {
	//all the handlers here

}
func main() {
	r := chi.NewRouter()
	app := &App{
		DB:     nil,
		Logger: nil,
	}
	app.mount()

	//if there's no need to use some global middleware on the mux, we can directly use http's package defaultMux (passing nil)

}
