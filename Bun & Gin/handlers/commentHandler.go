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
)

func CreateComment(c *gin.Context) {
	insertResult := insertComment(c.Request.Body)

	res := response{
		Message: insertResult,
	}

	c.JSON(http.StatusOK, res)
}

func GetAllComment(c *gin.Context) {

	all := getAllComment()

	c.JSON(http.StatusOK, all)
}

func GetComment(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	logFatal(err)

	comment, err := getComment(int64(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Detail not found",
		})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func UpdateComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	logFatal(err)

	updatedRows := updateComment(int64(id), c.Request.Body)

	msg := fmt.Sprintf("Comment updated successfully. Total rows/record affected %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	c.JSON(http.StatusOK, res)
}

func DeleteComment(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	logFatal(err)

	deletedRows := deleteComment(int64(id))

	msg := fmt.Sprintf("Comment deleted successfully. Total rows/record affected %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	c.JSON(http.StatusOK, res)
}

func insertComment(body io.ReadCloser) sql.Result {
	ctx := context.Background()
	var comment models.Comment

	err := json.NewDecoder(body).Decode(&comment)
	logFatal(err)

	db := createConnection()

	defer db.Close()

	res, err := db.NewInsert().Model(&comment).Exec(ctx)
	logFatal(err)

	return res
}

func getAllComment() []models.Comment {
	ctx := context.Background()
	var comments []models.Comment

	db := createConnection()

	defer db.Close()

	err := db.NewSelect().Model(&comments).OrderExpr("id ASC").Scan(ctx)
	logFatal(err)

	return comments
}

func getComment(id int64) (models.Comment, error) {
	ctx := context.Background()
	var comment models.Comment

	db := createConnection()

	defer db.Close()

	err := db.NewSelect().Model(&comment).Where("id = ?", id).Scan(ctx)
	if err != nil {
		log.Println(err)
		return models.Comment{}, err
	}

	return comment, nil
}

func updateComment(id int64, body io.ReadCloser) int64 {
	ctx := context.Background()
	var comment models.Comment

	err := json.NewDecoder(body).Decode(&comment)
	logFatal(err)

	db := createConnection()

	defer db.Close()

	res, err := db.NewUpdate().Model(&comment).Where("id = ?", id).Exec(ctx)
	logFatal(err)

	rowsAffected, err := res.RowsAffected()
	logFatal(err)

	return rowsAffected
}

func deleteComment(id int64) int64 {
	ctx := context.Background()

	db := createConnection()

	defer db.Close()

	res, err := db.NewDelete().Model((*models.Comment)(nil)).Where("id = ?", id).Exec(ctx)
	logFatal(err)

	rowsAffected, err := res.RowsAffected()
	logFatal(err)

	return rowsAffected
}
