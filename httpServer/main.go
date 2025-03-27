package main

import (
	"log"
	"net/http"
	"os"

	"github.com/E-nkv/backend-dev-projects/httpServer/api"
	"github.com/E-nkv/backend-dev-projects/httpServer/service"

	"github.com/go-chi/chi/v5"
)

var addr = "localhost:8080"

func main() {
	r := chi.NewRouter()
	app := &api.App{
		Service: &service.InMemoryService{},
		Log:     log.New(os.Stdout, "|| ", log.Ldate|log.Ltime|log.Lmsgprefix),
	}
	app.Mount(r)
	app.Log.Println("running the server at ", addr)
	log.Fatal(http.ListenAndServe(addr, r))

}
