package main

import (
	"fmt"
	"log"
	"net/http"
)

//this is what most tutorials look like. sure, u get to know that "http" package exists, and allows you to set up your server.
//but what structure should you use? is this even the way it's meant to be "done"? of course not.

// quick tutorial on how to set up a web server, yay! (trying to make it look dumb)
func main() {
	//route for when the user hits the home route (is this comment even needed? come on)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world!")
	})
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		//haha, will you really handle method based routing like this in a production app? how is the learner supposed to know this if "the tutorial" doesnt teach it???
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "getting all users!")
		case http.MethodPost:
			fmt.Fprint(w, "succesfully created a user!")
		default:
			fmt.Fprintf(w, "unknown method yay!") //sure bro, here the learner will surely understand advanced error handling from this default clause.
		}
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	//by the way, dear tutorial, didnt you know that u can do it inline? log.Fatal(http.ListenAndServe(":8080", nil)) ?? is there any difference
}
