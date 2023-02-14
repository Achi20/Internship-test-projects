package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go_test3/models"
	"io"
	"log"
	"net/http"
	"time"

	// "reflect"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "OSG1practice?"
	dbname   = "test3"
)

type response struct {
	ID      int64       `json:"id,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createConnection() *sqlx.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", psqlconn)
	logFatal(err)

	fmt.Println("Successfully connected!")
	return db
}

func Create2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")

	params := mux.Vars(r)

	insertResult := insert2(params["table"], r.Body)

	res := response{
		Message: insertResult,
	}

	json.NewEncoder(w).Encode(res)
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")

	params := mux.Vars(r)
	fmt.Println(params)

	switch params["table"] {
	case "user":
		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)
		logFatal(err)

		insertResult := insert(user, models.Post{}, models.Comment{})
		res := response{
			Message: insertResult,
		}

		json.NewEncoder(w).Encode(res)

	case "post":
		var post models.Post

		err := json.NewDecoder(r.Body).Decode(&post)
		logFatal(err)

		insertResult := insert(models.User{}, post, models.Comment{})
		res := response{
			Message: insertResult,
		}

		json.NewEncoder(w).Encode(res)

	case "comment":
		var comment models.Comment

		err := json.NewDecoder(r.Body).Decode(&comment)
		logFatal(err)

		insertResult := insert(models.User{}, models.Post{}, comment)
		res := response{
			Message: insertResult,
		}

		json.NewEncoder(w).Encode(res)
	}
	// table := define(params["table"]) //func grade() interface{} {
	// 	switch params["table"] {
	// 	case "user":
	// 		var user models.User
	// 		return user
	// 	case "post":
	// 		var post models.Post
	// 		return post
	// 	case "comment":
	// 		var comment models.Comment
	// 		return comment
	// 	default:
	// 		return nil
	// 	}
	// }
	// fmt.Println(110, table)
	// space
	// for i := range table {
	// space
	// }
	// space
	// err := json.NewDecoder(r.Body).Decode(&table)
	// logFatal(err)
	// fmt.Println(117, table)
	// space
	// insertResult := insertUser(table, params["table"])
	// space
	// res := response{
	// 	Message: insertResult,
	// }
	// space
	// json.NewEncoder(w).Encode(res)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	all := getAll(params["table"])

	json.NewEncoder(w).Encode(all)
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	logFatal(err)

	table := get(int64(id), params["table"])

	json.NewEncoder(w).Encode(table)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "PUT")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	logFatal(err)

	// var user models.User
	// err = json.NewDecoder(r.Body).Decode(&user)
	// logFatal(err)

	updatedRows := update(int64(id), params["table"], r.Body)

	msg := fmt.Sprintf("%v updated successfully. Total rows/record affected %v", params["table"], updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	logFatal(err)

	deletedUsers, deletedPosts, deletedComs := delete(int64(id), params["table"])

	msg := fmt.Sprintf("%v deleted successfully. Total users, posts, comments affected %v, %v, %v", params["table"], deletedUsers, deletedPosts, deletedComs)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func insert(user models.User, post models.Post, comment models.Comment) sql.Result { //table interface{}, name string) sql.Result {
	db := createConnection()

	defer db.Close()

	tx := db.MustBegin()

	switch {
	case user.Username != "":
		result := tx.MustExec("INSERT INTO users (username, password, avatar) VALUES ($1, $2, $3) RETURNING id", user.Username, user.Password, user.Avatar)
		err := tx.Commit()
		logFatal(err)
		return result

	case post.Description != "":
		result := tx.MustExec("INSERT INTO post (description, image, user_id, created_at) VALUES ($1, $2, $3, %4) RETURNING id", post.Description, post.Image, post.User_ID, time.Now())
		err := tx.Commit()
		logFatal(err)
		return result

	case comment.Text != "":
		result := tx.MustExec("INSERT INTO comment (post_id, user_id, text) VALUES ($1, $2, $3) RETURNING id", comment.Post_ID, comment.User_ID, comment.Text)
		err := tx.Commit()
		logFatal(err)
		return result
	}

	err := tx.Commit()
	logFatal(err)

	return nil
	// switch name {
	// case "user":
	// 	val := reflect.ValueOf(table)
	// 	for val.Next() {
	// 		k := val.Key()
	// 		if k == "ff" {
	// 			fmt.Println(3)
	// 		}
	// 		v := val.Value()
	// 	}
	// 	username := val.FieldByName("username").Interface().(string)
	// 	password := val.FieldByName("password").Interface().(string)
	// 	avatar := val.FieldByName("avatar").Interface().(string)
	// 	result := tx.MustExec("INSERT INTO users (username, password, avatar) VALUES ($1, $2, $3) RETURNING id", username, password, avatar)
	// 	err := tx.Commit()
	// 	logFatal(err)
	// space
	// 	return result
	// case "post":
	// 	result := tx.MustExec("INSERT INTO users (username, password, avatar) VALUES ($1, $2, $3) RETURNING id", user.Username, user.Password, user.Avatar)
	// 	err := tx.Commit()
	// 	logFatal(err)
	// space
	// 	return result
	// case "comment":
	// 	result := tx.MustExec("INSERT INTO users (username, password, avatar) VALUES ($1, $2, $3) RETURNING id", user.Username, user.Password, user.Avatar)
	// 	err := tx.Commit()
	// 	logFatal(err)
	// space
	// 	return result
	// }
	// return nil
	// space
	// result := tx.MustExec("INSERT INTO users (username, password, avatar) VALUES ($1, $2, $3) RETURNING id", user.Username, user.Password, user.Avatar)
	// err := tx.Commit()
	// logFatal(err)
	// return result
}

func insert2(name string, body io.ReadCloser) sql.Result {
	db := createConnection()

	defer db.Close()

	tx := db.MustBegin()

	switch name {
	case "user":
		var user models.User

		err := json.NewDecoder(body).Decode(&user)
		logFatal(err)

		result := tx.MustExec("INSERT INTO users (username, password, avatar) VALUES ($1, $2, $3) RETURNING id", user.Username, user.Password, user.Avatar)

		err = tx.Commit()
		logFatal(err)

		return result

	case "post":
		var post models.Post

		err := json.NewDecoder(body).Decode(&post)
		logFatal(err)

		result := tx.MustExec("INSERT INTO posts (description, image, user_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id", post.Description, post.Image, post.User_ID, time.Now())

		err = tx.Commit()
		logFatal(err)

		return result

	case "comment":
		var comment models.Comment

		err := json.NewDecoder(body).Decode(&comment)
		logFatal(err)

		result := tx.MustExec("INSERT INTO comments (post_id, user_id, text) VALUES ($1, $2, $3) RETURNING id", comment.Post_ID, comment.User_ID, comment.Text)

		err = tx.Commit()
		logFatal(err)

		return result
	}

	err := tx.Commit()
	logFatal(err)

	return nil
}

func getAll(name string) interface{} {
	db := createConnection()

	defer db.Close()

	switch name {
	case "user":
		var users []models.User
		err := db.Select(&users, "SELECT * FROM users")
		logFatal(err)
		return users

	case "post":
		var posts []models.Post
		err := db.Select(&posts, "SELECT * FROM posts")
		logFatal(err)
		return posts

	case "comment":
		var comments []models.Comment
		err := db.Select(&comments, "SELECT * FROM comments")
		logFatal(err)
		return comments
	}

	return nil
	// var users []models.User
	// db := createConnection()
	// defer db.Close()
	// err := db.Select(&users, "SELECT * FROM users")
	// logFatal(err)
	// return users
}

func get(id int64, name string) interface{} {
	db := createConnection()

	defer db.Close()

	switch name {
	case "user":
		var user models.User
		err := db.Get(&user, "SELECT * FROM users WHERE id=$1", id)
		logFatal(err)
		return user

	case "post":
		var post models.Post
		err := db.Get(&post, "SELECT * FROM posts WHERE id=$1", id)
		logFatal(err)
		return post

	case "comment":
		var comment models.Comment
		err := db.Get(&comment, "SELECT * FROM comments WHERE id=$1", id)
		logFatal(err)
		return comment
	}

	return nil
	// db := createConnection()
	// var user models.User
	// defer db.Close()
	// err := db.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	// logFatal(err)
	// return user
}

func update(id int64, name string, body io.ReadCloser) int64 {
	db := createConnection()

	defer db.Close()

	tx := db.MustBegin()

	switch name {
	case "user":
		var user models.User

		err := json.NewDecoder(body).Decode(&user)
		logFatal(err)

		result := tx.MustExec("UPDATE users SET username=$2, password=$3, avatar=$4 WHERE id=$1", id, user.Username, user.Password, user.Avatar)

		err = tx.Commit()
		logFatal(err)

		rowsAffected, err := result.RowsAffected()
		logFatal(err)

		return rowsAffected

	case "post":
		var post models.Post

		err := json.NewDecoder(body).Decode(&post)
		logFatal(err)

		result := tx.MustExec("UPDATE posts SET description=$2, image=$3, user_id=$4 WHERE id=$1", id, post.Description, post.Image, post.User_ID)

		err = tx.Commit()
		logFatal(err)

		rowsAffected, err := result.RowsAffected()
		logFatal(err)

		return rowsAffected

	case "comment":
		var comment models.Comment

		err := json.NewDecoder(body).Decode(&comment)
		logFatal(err)

		result := tx.MustExec("UPDATE comments SET post_id=$2, user_id=$3, text=$4 WHERE id=$1", id, comment.Post_ID, comment.User_ID, comment.Text)

		err = tx.Commit()
		logFatal(err)

		rowsAffected, err := result.RowsAffected()
		logFatal(err)

		return rowsAffected
	}

	err := tx.Commit()
	logFatal(err)

	return 0
	// db := createConnection()
	// defer db.Close()
	// tx := db.MustBegin()
	// result := tx.MustExec("UPDATE users SET username=$2, password=$3, avatar=$4 WHERE id=$1", id, user.Username, user.Password, user.Avatar)
	// err := tx.Commit()
	// logFatal(err)
	// rowsAffected, err := result.RowsAffected()
	// logFatal(err)
	// return rowsAffected
}

func delete(id int64, name string) (int64, int64, int64) {
	db := createConnection()

	defer db.Close()

	tx := db.MustBegin()

	switch name {
	case "user":
		comment := tx.MustExec("DELETE FROM comments WHERE user_id=$1", id)
		post := tx.MustExec("DELETE FROM posts WHERE user_id=$1", id)
		result := tx.MustExec("DELETE FROM users WHERE id=$1", id)

		err := tx.Commit()
		logFatal(err)

		rowsAffected, err := result.RowsAffected()
		logFatal(err)

		postsAffected, err := post.RowsAffected()
		logFatal(err)

		comsAffected, err := comment.RowsAffected()
		logFatal(err)

		return rowsAffected, postsAffected, comsAffected

	case "post":
		comment := tx.MustExec("DELETE FROM comments WHERE post_id=$1", id)
		result := tx.MustExec("DELETE FROM posts WHERE id=$1", id)
		err := tx.Commit()
		logFatal(err)

		rowsAffected, err := result.RowsAffected()
		logFatal(err)

		comsAffected, err := comment.RowsAffected()
		logFatal(err)

		return 0, rowsAffected, comsAffected

	case "comment":
		result := tx.MustExec("DELETE FROM comments WHERE id=$1", id)
		err := tx.Commit()
		logFatal(err)

		rowsAffected, err := result.RowsAffected()
		logFatal(err)

		return 0, 0, rowsAffected
	}

	err := tx.Commit()
	logFatal(err)

	return 0, 0, 0
	// db := createConnection()
	// defer db.Close()
	// tx := db.MustBegin()
	// result := tx.MustExec("DELETE FROM users WHERE id=$1", id)
	// err := tx.Commit()
	// logFatal(err)
	// rowsAffected, err := result.RowsAffected()
	// logFatal(err)
	// return rowsAffected
}

// func define(name string) interface{} {
// 	fmt.Println(272, name)
// 	switch name {
// 	case "user":
// 		var user models.User
// 		fmt.Println(278, user)
// 		return user
// 	case "post":
// 		var post models.Post
// 		return post
// 	case "comment":
// 		var comment models.Comment
// 		return comment
// 	default:
// 		return nil
// 	}
// }
