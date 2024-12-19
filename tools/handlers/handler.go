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
	todo := router.Group("/")
	{
		todo.GET("/", func(c *gin.Context) {
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
		todo.GET("/signup", func(c *gin.Context) {
			HGET_signup(c)
		})
		todo.POST("/signup", func(c *gin.Context) {
			HPOST_signup(c, storage)
		})
		todo.GET("/login", func(c *gin.Context) {
			HGET_login(c)
		})
		todo.POST("/login", func(c *gin.Context) {
			HPOST_login(c, storage)
		})
		todo.POST("/logout", HPOST_logout)
	}
	return &Router{
		Router: router,
		Storage: storage,
	}
}
