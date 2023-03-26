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

var (
	categories = make([]Category, 0)
	bookes     = make([]Book, 0)
	connect    = make([]Connect, 0)
	b          Book
	c          Category
	conn       Connect
)

func GetBookes() []Book {

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

// func (m *Book) Create() *Book {
// 	//director bor yoki yo'qligini tekshirish uchun
// 	row, err := db.Query("SELECT id FROM director WHERE lastname = ? and firstname = ?", m.Director.Lastname, m.Director.Firstname)
// 	var id int
// 	for row.Next() {
// 		err = row.Scan(&id)
// 		if err != nil {
// 			log.Fatal(err)
// 			os.Exit(1)
// 		}
// 	}
// 	defer row.Close()
// 	if id == 0 { //agar director mavjud bo'lmasa uni yaratamiz
// 		row, err := db.Query("INSERT INTO `director`(`firstname`, `lastname`) VALUES(?,?)", m.Director.Firstname, m.Director.Lastname)
// 		if err != nil {
// 			log.Fatal(err)
// 			os.Exit(1)
// 		}
// 		defer row.Close()
// 	}
// 	// director id ni olamiz
// 	row, err = db.Query("SELECT id FROM director WHERE lastname = ? and firstname = ?", m.Director.Lastname, m.Director.Firstname)
// 	for row.Next() {
// 		err = row.Scan(&id)
// 		if err != nil {
// 			log.Fatal(err)
// 			os.Exit(1)
// 		}
// 	}
// 	m.DirectorId = id
// 	res, err := db.Query("INSERT INTO `movie` (`created_at`, `isbn`, `title`, `directorId`) VALUES (NOW(), ?, ?, ?)", m.Isbn, m.Title, id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer res.Close()
// 	return m
// }

// func Delete(id int) *Movie {
// 	movie := GetMovie(id)
// 	row, err := db.Query("DELETE FROM movie WHERE movieId = ?", id)
// 	if err != nil {
// 		log.Fatal(err)
// 		os.Exit(1)
// 	}
// 	defer row.Close()
// 	return movie
// }

// func (m *Movie) Update(id int) *Movie {
// 	// get directorId
// 	directoId := 0
// 	row, err := db.Query("SELECT id FROM director WHERE lastname = ? and firstname = ?", m.Director.Lastname, m.Director.Firstname)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for row.Next() {
// 		err = row.Scan(&directoId)
// 		if err != nil {
// 			log.Fatal(err)
// 			os.Exit(1)
// 		}
// 	}
// 	defer row.Close()
// 	if directoId == 0 {
// 		//agar director mavjud bo'lmasa uni yaratamiz
// 		row, err := db.Query("INSERT INTO `director`(`firstname`, `lastname`) VALUES(?,?)", m.Director.Firstname, m.Director.Lastname)
// 		if err != nil {
// 			log.Fatal(err)
// 			os.Exit(1)
// 		}
// 		defer row.Close()
// 	}
// 	// director id ni olamiz
// 	row, err = db.Query("SELECT id FROM director WHERE lastname = ? and firstname = ?", m.Director.Lastname, m.Director.Firstname)
// 	for row.Next() {
// 		err = row.Scan(&directoId)
// 		if err != nil {
// 			log.Fatal(err)
// 			os.Exit(1)
// 		}
// 	}

// 	// yangilash update
// 	row, err = db.Query("UPDATE movie SET title = ?, isbn = ?, directorId = ? where movieId = ?", m.Title, m.Isbn, directoId, id)
// 	if err != nil {
// 		return nil
// 	}
// 	defer row.Close()
// 	return m
// }
