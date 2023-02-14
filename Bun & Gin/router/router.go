package router

import (
	"go_test5/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/user", handlers.GetAllUser)
	router.GET("/user/:id", handlers.GetUser)
	router.POST("/new/user", handlers.CreateUser)
	router.PUT("/user/:id", handlers.UpdateUser)
	router.DELETE("/delete/user/:id", handlers.DeleteUser)

	router.GET("/post", handlers.GetAllPost)
	router.GET("/post/:id", handlers.GetPost)
	router.POST("/new/post", handlers.CreatePost)
	router.PUT("/post/:id", handlers.UpdatePost)
	router.DELETE("/delete/post/:id", handlers.DeletePost)

	router.GET("/comment", handlers.GetAllComment)
	router.GET("/comment/:id", handlers.GetComment)
	router.POST("/new/comment", handlers.CreateComment)
	router.PUT("/comment/:id", handlers.UpdateComment)
	router.DELETE("/delete/comment/:id", handlers.DeleteComment)

	return router
}
