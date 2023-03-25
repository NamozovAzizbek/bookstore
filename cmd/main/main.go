package main

import (
	"log"
	"net/http"

	routers "github.com/NamozovAzizbek/bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routers.RegisterMovieRouters(r)
	r.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
