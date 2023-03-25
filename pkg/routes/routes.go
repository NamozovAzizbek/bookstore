package routers

import(
	"github.com/NamozovAzizbek/bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterMovieRouters = func (router *mux.Router)  {
	router.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	router.HandleFunc("/movie/{id}", controllers.GetMovieById).Methods("GET")
	router.HandleFunc("/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movie/{id}", controllers.UpdateMovie).Methods("PUT")	
	router.HandleFunc("/movie/{id}", controllers.DeleteMovie).Methods("DELETE")

}