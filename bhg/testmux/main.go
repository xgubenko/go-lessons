package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type trivial struct{}

func (t *trivial) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("Executing trivial middleware")
	next(w, r)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users/{user}", func(w http.ResponseWriter, req *http.Request) {
		user := mux.Vars(req)["user"]
		fmt.Fprintf(w, "hi %s\n", user)
	}).Methods("GET")
	n := negroni.Classic()
	n.UseHandler(r)
	n.Use(&trivial{})
	http.ListenAndServe(":8080", n)
}
