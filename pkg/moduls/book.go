package moduls

import (
	"log"
	"os"

	"github.com/NamozovAzizbek/bookstore/pkg/config"
)

type Category struct {
	Id            int    `json:"id"`
	Book_category string `json:"book_category"`
}

type Book struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	Category    Category
}

type Connect struct {
	Id_book     int
	Id_Category int
}

var db = config.Connect()

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func GetBookes() []Book {

	var (
		categories = make([]Category, 0)
		bookes     = make([]Book, 0)
		connect    = make([]Connect, 0)
		b          Book
		c          Category
		conn       Connect
	)
	row, err := db.Query("SELECT * from category")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer row.Close()

	for row.Next() {
		err := row.Scan(&c.Id, &c.Book_category)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		categories = append(categories, c)
	}

	row, err = db.Query("SELECT * from connect")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer row.Close()

	for row.Next() {
		err := row.Scan(&conn.Id_book, &conn.Id_Category)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		connect = append(connect, conn)
	}

	row, err = db.Query("SELECT * from book")
	if err != nil {
		log.Fatalf("book query:%v", err)
		return nil
	}
	defer row.Close()

	for row.Next() {
		err := row.Scan(&b.ID, &b.Name, &b.Author, &b.Publication)
		if err != nil {
			log.Fatalf("book Scan:%v", err)
			return nil
		}

		for _, v := range connect {
			if v.Id_book == b.ID {
				for _, cat := range categories {
					if cat.Id == v.Id_Category {
						b.Category = cat
						bookes = append(bookes, b)
					}
				}
			}
		}
	}
	return bookes
}

func GetBook(id int) []Book {

	var (
		categories = make([]Category, 0)
		bookes     = make([]Book, 0)
		connect    = make([]Connect, 0)
		b          Book
		c          Category
		conn       Connect
	)

	row, err := db.Query("SELECT * from category")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer row.Close()

	for row.Next() {
		err := row.Scan(&c.Id, &c.Book_category)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		categories = append(categories, c)
	}

	row, err = db.Query("SELECT * from connect")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer row.Close()

	for row.Next() {
		err := row.Scan(&conn.Id_book, &conn.Id_Category)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		connect = append(connect, conn)
	}

	row, err = db.Query("SELECT * from book where id = $1 ", id)
	if err != nil {
		log.Fatalf("book by id query:%v", err)
		return nil
	}

	defer row.Close()

	for row.Next() {
		err := row.Scan(&b.ID, &b.Name, &b.Author, &b.Publication)
		if err != nil {
			log.Fatalf("book by id scan:%v", err)
			return nil
		}
	}
	for _, v := range connect {
		if v.Id_book == b.ID {
			for _, cat := range categories {
				if cat.Id == v.Id_Category {
					b.Category = cat
					bookes = append(bookes, b)
				}
			}
		}
	}

	return bookes
}

func (b *Book) Create() *Book {
	var book Book
	// categoryani aniqlash
	row, err := db.Query("select id from category where category=$1", b.Category.Book_category)
	checkError(err)
	for row.Next() {
		err = row.Scan(&book.Category.Id)
		checkError(err)
	}
	defer row.Close()
	if book.Category.Id == 0 {
		_, err = db.Exec("insert into category(category) values ($1)", b.Category.Book_category)
		checkError(err)
	}
	// category id olindi
	row, err = db.Query("select id from category where category = $1", b.Category.Book_category)
	checkError(err)
	for row.Next() {
		err = row.Scan(&book.Category.Id)
		checkError(err)
	}

	_, err = db.Exec("insert into book(name, author, publication) values ($1, $2, $3)", b.Name, b.Author, b.Publication)
	checkError(err)

	return b
}

func Delete(id int) {

	_, err := db.Exec("DELETE FROM book WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	
}

func (b *Book) Update(id int) *Book {
	// get id book
	bookId := 0
	row, err := db.Query("SELECT id FROM book WHERE name = $1 and author = $2 and publication = $3", b.Name, b.Author, b.Publication)
	if err != nil {
		log.Fatal(err)
	}
	for row.Next(){
		err = row.Scan(&bookId)
		checkError(err)
	}
	defer row.Close()
	if bookId == 0{
		return nil
	}
	// yangilash update
	row, err = db.Query("UPDATE book SET name = 1$, author = $2, publication = $3 where id = ?", b.Name, b.Author, b.Publication, id)
	checkError(err)
	defer row.Close()
	return b
}
