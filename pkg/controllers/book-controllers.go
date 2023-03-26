package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NamozovAzizbek/bookstore/pkg/moduls"
	"github.com/gorilla/mux"
)

var NewBook moduls.Book

func GetBookes(w http.ResponseWriter, r *http.Request) {
	bookes := moduls.GetBookes()
	res, _ := json.Marshal(bookes)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	req := mux.Vars(r)["id"]
	id, err := strconv.Atoi(req)

	if err != nil {
        w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid product ID")
        return
    }
	book := moduls.GetBook(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// func CreateMovie(w http.ResponseWriter, r *http.Request) {
// 	newMovie := &moduls.Movie{}
// 	utils.ParseBody(r, &newMovie)
// 	newMovie.Create()
// 	res, _ := json.Marshal(newMovie)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func DeleteMovie(w http.ResponseWriter, r *http.Request) {
// 	req := mux.Vars(r)["id"]
// 	id, err := strconv.Atoi(req)
// 	if err != nil {
// 		fmt.Println("errror while parsing")
// 		return
// 	}
// 	movie := moduls.Delete(id)

// 	res, _ := json.Marshal(movie)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func UpdateMovie(w http.ResponseWriter, r *http.Request) {
// 	req := mux.Vars(r)["id"]
// 	id, err := strconv.Atoi(req)
// 	if err != nil {
// 		fmt.Println("errror while parsing")
// 		return
// 	}
// 	movie := &moduls.Movie{}
// 	utils.ParseBody(r, &movie)
// 	movie.Update(id)

// 	if movie.Update(id) == nil {
// 		fmt.Fprintf(w, "movie not found")
// 	} else {
// 		res, _ := json.Marshal(movie)
// 		w.Header().Set("Content-Type", "pkglication/json")
// 		w.WriteHeader(http.StatusOK)
// 		w.Write(res)
// 	}
// }
