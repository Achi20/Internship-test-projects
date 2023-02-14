package controllers

import (
	"database/sql"
	"encoding/json"
	"go_practice_http/models"
	bookRepository "go_practice_http/repository/book"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Controller struct{}

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}

		books = bookRepo.GetBooks(db, book, books)
		// rows, err := db.Query("select * from books")
		// logFatal(err)
		// defer rows.Close()
		// for rows.Next() {
		// 	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		// 	logFatal(err)
		// 	books = append(books, book)
		// }
		json.NewEncoder(w).Encode(books)
	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		params := mux.Vars(r)
		bookRepo := bookRepository.BookRepository{}

		id, err := strconv.Atoi(params["id"])
		logFatal(err)

		book = bookRepo.GetBook(db, book, id)
		// rows := db.QueryRow("select * from books where id=$1", params["id"])
		// err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		// logFatal(err)
		json.NewEncoder(w).Encode(book)
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		bookRepo := bookRepository.BookRepository{}

		json.NewDecoder(r.Body).Decode(&book)

		bookID := bookRepo.AddBook(db, book)
		// err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;",
		// 	book.Title, book.Author, book.Year).Scan(&bookID)
		// logFatal(err)
		json.NewEncoder(w).Encode(bookID)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		bookRepo := bookRepository.BookRepository{}

		json.NewDecoder(r.Body).Decode(&book)

		rowsUpdated := bookRepo.UpdateBook(db, book)
		// result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id",
		// 	book.Title, book.Author, book.Year, book.ID)
		// logFatal(err)
		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		bookRepo := bookRepository.BookRepository{}

		id, err := strconv.Atoi(params["id"])
		logFatal(err)

		rowsDeleted := bookRepo.RemoveBook(db, id)
		// result, err := db.Exec("delete from books where id = $1", params["id"])
		// logFatal(err)
		json.NewEncoder(w).Encode(rowsDeleted)
	}
}
