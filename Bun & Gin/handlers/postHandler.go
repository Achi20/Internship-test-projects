package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go_test5/models"
	"io"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	insertResult := insertPost(c.Request.Body)

	res := response{
		Message: insertResult,
	}

	c.JSON(http.StatusOK, res)
}

func GetAllPost(c *gin.Context) {
	all := getAllPost()

	c.JSON(http.StatusOK, all)
}

func GetPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	logFatal(err)

	post := getPost(int64(id))

	c.JSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	logFatal(err)

	updatedRows := updatePost(int64(id), c.Request.Body)

	msg := fmt.Sprintf("Post updated successfully. Total rows/record affected %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	c.JSON(http.StatusOK, res)
}

func DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	logFatal(err)

	deletedPosts, deletedComms := deletePost(int64(id))

	msg := fmt.Sprintf("Post deleted successfully. Total posts and comments affected: %v, %v", deletedPosts, deletedComms)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	c.JSON(http.StatusOK, res)
}

func insertPost(body io.ReadCloser) sql.Result {
	ctx := context.Background()
	var post models.Post

	err := json.NewDecoder(body).Decode(&post)
	logFatal(err)

	db := createConnection()

	defer db.Close()
	fmt.Println(post)
	res, err := db.NewInsert().Model(&post).Exec(ctx)
	logFatal(err)

	return res
}

func getAllPost() []models.Post {
	ctx := context.Background()
	var posts []models.Post

	db := createConnection()

	defer db.Close()

	err := db.NewSelect().Model(&posts).OrderExpr("id ASC").Scan(ctx)
	logFatal(err)

	return posts
}

func getPost(id int64) models.Post {
	ctx := context.Background()
	var post models.Post

	db := createConnection()

	defer db.Close()

	err := db.NewSelect().Model(&post).Where("id = ?", id).Scan(ctx)
	logFatal(err)

	return post
}

func updatePost(id int64, body io.ReadCloser) int64 {
	ctx := context.Background()
	var post models.Post

	err := json.NewDecoder(body).Decode(&post)
	logFatal(err)

	db := createConnection()

	defer db.Close()

	res, err := db.NewUpdate().Model(&post).ExcludeColumn("created_at").Where("id = ?", id).Exec(ctx)
	logFatal(err)

	rowsAffected, err := res.RowsAffected()
	logFatal(err)

	return rowsAffected
}

func deletePost(id int64) (int64, int64) {
	ctx := context.Background()

	db := createConnection()

	defer db.Close()

	comment, err := db.NewDelete().Model((*models.Comment)(nil)).Where("post_id = ?", id).Exec(ctx)
	logFatal(err)

	post, err := db.NewDelete().Model((*models.Post)(nil)).Where("id = ?", id).Exec(ctx)
	logFatal(err)

	postsAffected, err := post.RowsAffected()
	logFatal(err)

	commsAffected, err := comment.RowsAffected()
	logFatal(err)

	return postsAffected, commsAffected
}
