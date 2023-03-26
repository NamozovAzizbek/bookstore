package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NamozovAzizbek/bookstore/pkg/moduls"
	"github.com/NamozovAzizbek/bookstore/pkg/utils"
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
	if len(book) == 0 {
		fmt.Fprintf(w, "book not found !")
		return
	}
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &moduls.Book{}
	utils.ParseBody(r, &newBook)
	newBook.Create()
	res, _ := json.Marshal(newBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	req := mux.Vars(r)["id"]
	id, err := strconv.Atoi(req)
	if err != nil {
		fmt.Println("errror while parsing")
		return
	}
	moduls.Delete(id)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "deleted successfully")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	req := mux.Vars(r)["id"]
	id, err := strconv.Atoi(req)
	if err != nil {
		fmt.Println("errror while parsing")
		return
	}
	book := &moduls.Book{}
	utils.ParseBody(r, &book)
	book.Update(id)

	if book.Update(id) == nil {
		fmt.Fprintf(w, "book not found")
		return
	} else {
		res, _ := json.Marshal(book)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
