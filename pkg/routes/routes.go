package routers

import(
	"github.com/NamozovAzizbek/bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookRouters = func (router *mux.Router)  {
	router.HandleFunc("/books", controllers.GetBookes).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")	
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")

}