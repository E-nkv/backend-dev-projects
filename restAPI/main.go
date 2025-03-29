package main

import (
	"log"
	"net/http"
	"os"

	"github.com/E-nkv/backend-dev-projects/restAPI/api"
	"github.com/E-nkv/backend-dev-projects/restAPI/service"

	"github.com/go-chi/chi/v5"
)

var addr = "localhost:8080"

func main() {
	l := log.New(os.Stdout, "|| ", log.Ldate|log.Ltime|log.Lmsgprefix)
	l.Println("welcome to restAPI")
	r := chi.NewRouter()
	app := &api.App{
		Service: &service.InMemoryService{},
		Log:     l,
	}
	app.Mount(r)
	app.Log.Println("running the server at ", addr)
	log.Fatal(http.ListenAndServe(addr, r))

}
