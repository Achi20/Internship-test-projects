package main

import (
	// "database/sql"
	"fmt"
	"go_test3/router"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "OSG1practice?"
// 	dbname   = "test3"
// )

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 8080...")
	logFatal(http.ListenAndServe(":8080", r))
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	// db, err := sql.Open("postgres", psqlconn)
	// logFatal(err)
	// defer db.Close()
	// err = db.Ping()
	// logFatal(err)
	// insertStmt := `insert into "users"("username", "password") values('John', '0faf')`
	// _, err = db.Exec(insertStmt)
	// logFatal(err)
	// insertDynStmt := `insert into "users"("username", "password") values($1, $2)`
	// _, err = db.Exec(insertDynStmt, "Jane", "5iqieq")
	// logFatal(err)
	// updateStmt := `update "users" set "username"=$1, "password"=$2 where "id"=$3`
	// _, err = db.Exec(updateStmt, "Mary", "8jpqcd", 1)
	// logFatal(err)
	// deleteStmt := `delete from "users" where id=$1`
	// _, err = db.Exec(deleteStmt, 2)
	// logFatal(err)
	// rows, err := db.Query(`SELECT "username" FROM "users"`)
	// logFatal(err)
	// defer rows.Close()
	// for rows.Next() {
	// 	var name string
	// 	err = rows.Scan(&name)
	// 	logFatal(err)
	// 	fmt.Println(name)
	// }
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
