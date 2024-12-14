package handlers

import (
	"pet_pr/tools/storage"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Router *gin.Engine
	Storage *storage.Storage
}

func InitHandlers(storage *storage.Storage) *Router {
	router := gin.New()
	router.LoadHTMLGlob("./tools/handlers/htmls/*")
	user := router.Group("/user")
	{
		user.POST("/", PostUser)
		user.GET("/", GetUsers)
		user.GET("/id=:id", func(c *gin.Context) {
			GetUserById(c, storage)
		})
		user.POST("/", func(c *gin.Context) {
			PushUser(c, storage)
		})
	}
	todo := router.Group("/todo")
	{
		todo.GET("/id=:id", func(c *gin.Context) {
			HandleUserTodo(c, storage)
		})
		todo.POST("/update_task", func(c *gin.Context) {
			HandleAnswerTodo(c, storage)
		})
		todo.POST("/add_task", func(c *gin.Context) {
			HandlePushTodo(c, storage)
		})
		todo.POST("/delete_task", func(c *gin.Context) {
			HandleDeleteTodo(c, storage)
		})
	}
	return &Router{
		Router: router,
		Storage: storage,
	}
}
