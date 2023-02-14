package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go_test5/models"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
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

func createConnection() *bun.DB {
	sqldb, err := sql.Open("postgres", "user=postgres password=OSG1practice? dbname=test3 sslmode=disable")
	logFatal(err)

	sqldb.SetMaxOpenConns(1)

	db := bun.NewDB(sqldb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	fmt.Println("Successfully connected!")
	return db
}

func CreateUser(c *gin.Context) {
	insertResult := insertUser(c.Request.Body)

	res := response{
		Message: insertResult,
	}

	c.JSON(http.StatusOK, res)
}

func GetAllUser(c *gin.Context) {
	all := getAllUser()

	c.JSON(http.StatusOK, all)
}

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	logFatal(err)

	user := getUser(int64(id))

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	logFatal(err)

	updatedRows := updateUser(int64(id), c.Request.Body)

	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	c.JSON(http.StatusOK, res)
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	logFatal(err)

	deletedUsers, deletedPosts, deletedComms := deleteUser(int64(id))

	msg := fmt.Sprintf("User deleted successfully. Total users, posts, comments affected: %v, %v, %v", deletedUsers, deletedPosts, deletedComms)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	c.JSON(http.StatusOK, res)
}

func insertUser(body io.ReadCloser) sql.Result {
	ctx := context.Background()
	var user models.User

	err := json.NewDecoder(body).Decode(&user)
	logFatal(err)

	db := createConnection()

	defer db.Close()

	res, err := db.NewInsert().Model(&user).Exec(ctx)
	logFatal(err)

	return res
}

func getAllUser() []models.User {
	ctx := context.Background()
	var users []models.User

	db := createConnection()

	defer db.Close()

	err := db.NewSelect().Model(&users).OrderExpr("id ASC").Scan(ctx)
	logFatal(err)

	return users
}

func getUser(id int64) models.User {
	ctx := context.Background()
	var user models.User

	db := createConnection()

	defer db.Close()

	err := db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx)
	logFatal(err)

	return user
}

func updateUser(id int64, body io.ReadCloser) int64 {
	ctx := context.Background()
	var user models.User

	err := json.NewDecoder(body).Decode(&user)
	logFatal(err)

	db := createConnection()

	defer db.Close()

	res, err := db.NewUpdate().Model(&user).Where("id = ?", id).Exec(ctx)
	logFatal(err)

	rowsAffected, err := res.RowsAffected()
	logFatal(err)

	return rowsAffected
}

func deleteUser(id int64) (int64, int64, int64) {
	ctx := context.Background()

	db := createConnection()

	defer db.Close()

	comment, err := db.NewDelete().Model((*models.Comment)(nil)).Where("user_id = ?", id).Exec(ctx)
	logFatal(err)

	cp := db.NewSelect().Model((*models.Post)(nil)).Where("user_id = ?", id)
	commentOfPost, err := db.NewDelete().Model((*models.Comment)(nil)).TableExpr("(?) AS comm", cp).Where("comment.post_id = comm.id").Exec(ctx)
	logFatal(err)

	post, err := db.NewDelete().Model((*models.Post)(nil)).Where("user_id = ?", id).Exec(ctx)
	logFatal(err)

	user, err := db.NewDelete().Model((*models.User)(nil)).Where("id = ?", id).Exec(ctx)
	logFatal(err)

	usersAffected, err := user.RowsAffected()
	logFatal(err)

	postsAffected, err := post.RowsAffected()
	logFatal(err)

	commsAffected, err := comment.RowsAffected()
	logFatal(err)

	commsPostsAffected, err := commentOfPost.RowsAffected()
	logFatal(err)

	return usersAffected, postsAffected, commsAffected + commsPostsAffected
}
