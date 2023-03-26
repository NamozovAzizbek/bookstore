package routers

import(
	"github.com/NamozovAzizbek/bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookRouters = func (router *mux.Router)  {
	router.HandleFunc("/books", controllers.GetBookes).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	// router.HandleFunc("/movie", controllers.CreateMovie).Methods("POST")
	// router.HandleFunc("/movie/{id}", controllers.UpdateMovie).Methods("PUT")	
	// router.HandleFunc("/movie/{id}", controllers.DeleteMovie).Methods("DELETE")

}