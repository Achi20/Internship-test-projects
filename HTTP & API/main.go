package main

import (
	"database/sql"
	"go_practice_http/controllers"
	"go_practice_http/driver"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	// "go_practice_http/models"
	// "encoding/json"
	// "os"
	// "github.com/lib/pq"
)

// type Book struct {
// 	ID     int    `json:id`
// 	Title  string `json:title`
// 	Author string `json: author`
// 	Year   string `json: year`
// }

// var books []models.Book
var db *sql.DB

func init() {
	gotenv.Load()
}

// func logFatal(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {
	// pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	// logFatal(err)
	// log.Println(pgUrl)
	// db, err = sql.Open("postgres", pgUrl)
	// logFatal(err)
	// err = db.Ping()
	// logFatal(err)
	db = driver.ConnectDB()
	router := mux.NewRouter()

	controller := controllers.Controller{}

	// books = append(books, Book{1, "Golang pointers", "Mr. Golang", "2010"},
	// 	Book{2, "Goroutines", "Mr. Goroutine", "2011"},
	// 	Book{3, "Golang routers", "Mr. Router", "2012"},
	// 	Book{4, "Golang concurrency", "Mr. Currency", "2013"},
	// 	Book{5, "Golang good parts", "Mr. Good", "2014"})

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// func getBooks(w http.ResponseWriter, r *http.Request) {
// 	var book models.Book
// 	books = []models.Book{}
// 	rows, err := db.Query("select * from books")
// 	logFatal(err)
// 	defer rows.Close()
// 	for rows.Next() {
// 		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
// 		logFatal(err)
// 		books = append(books, book)
// 	}
// 	json.NewEncoder(w).Encode(books)
// }

// func getBook(w http.ResponseWriter, r *http.Request) {
// 	var book models.Book
// 	params := mux.Vars(r)
// 	rows := db.QueryRow("select * from books where id=$1", params["id"])
// 	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
// 	logFatal(err)
// 	json.NewEncoder(w).Encode(book)
// 	// params := mux.Vars(r)
// 	// i, _ := strconv.Atoi(params["id"])
// 	// for _, book := range books {
// 	// 	if book.ID == i {
// 	// 		json.NewEncoder(w).Encode(&book)
// 	// 	}
// 	// }
// }

// func addBook(w http.ResponseWriter, r *http.Request) {
// 	var book models.Book
// 	var bookID int
// 	json.NewDecoder(r.Body).Decode(&book)
// 	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;",
// 		book.Title, book.Author, book.Year).Scan(&bookID)
// 	logFatal(err)
// 	json.NewEncoder(w).Encode(bookID)
// 	// var newBook Book
// 	// json.NewDecoder(r.Body).Decode(&newBook)
// 	// books = append(books, newBook)
// 	// json.NewEncoder(w).Encode(books)
// }

// func updateBook(w http.ResponseWriter, r *http.Request) {
// 	var book models.Book
// 	json.NewDecoder(r.Body).Decode(&book)
// 	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id",
// 		book.Title, book.Author, book.Year, book.ID)
// 	logFatal(err)
// 	rowsUpdated, err := result.RowsAffected()
// 	logFatal(err)
// 	json.NewEncoder(w).Encode(rowsUpdated)
// 	// var book Book
// 	// json.NewDecoder(r.Body).Decode(&book)
// 	// for i, item := range books {
// 	// 	if item.ID == book.ID {
// 	// 		books[i] = book
// 	// 	}
// 	// }
// 	// json.NewEncoder(w).Encode(books)
// }

// func removeBook(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	result, err := db.Exec("delete from books where id = $1", params["id"])
// 	logFatal(err)
// 	rowsDeleted, err := result.RowsAffected()
// 	logFatal(err)
// 	json.NewEncoder(w).Encode(rowsDeleted)
// 	// params := mux.Vars(r)
// 	// id, _ := strconv.Atoi(params["id"])
// 	// for i, item := range books {
// 	// 	if item.ID == id {
// 	// 		books = append(books[:i], books[i+1:]...)
// 	// 	}
// 	// }
// 	// json.NewEncoder(w).Encode(books)
// }
