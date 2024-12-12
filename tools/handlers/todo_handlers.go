package handlers

import (
	"fmt"
	"net/http"
	"pet_pr/tools/models"
	"pet_pr/tools/storage"

	"github.com/gin-gonic/gin"
)

func HandleUserTodo(c *gin.Context, storage *storage.Storage) {
	id := c.Param("id")
	todos := storage.GetUserTodosByUID(id)
	c.HTML(http.StatusOK, "todo.tmpl", gin.H {
		"Todos": todos,
 	})
}

func HandleAnswerTodo(c *gin.Context, storage *storage.Storage) {
	var answer models.TodoItem
	if err := c.BindJSON(&answer); err != nil {
		storage.Logger.Info(fmt.Sprintf("Cannot parse answer: %s", err.Error()))
		c.JSON(http.StatusBadRequest, "Not update")
		return
	}
	storage.UpdateDoneTask(answer.Done, storage.GetIDByTitle(answer.ID))
	storage.Logger.Info("Task updated")
}

func HandlePushTodo(c *gin.Context, storage *storage.Storage) {
	type Req struct {
		Title string
		UID string
	}
	var answ Req
	if err := c.BindJSON(&answ); err != nil {
		storage.Logger.Info(fmt.Sprintf("Cannot parse answer: %s", err.Error()))
		c.JSON(http.StatusBadRequest, "Not update")
		return
	}
	storage.PushTodoByUID(answ.Title, answ.UID)
	storage.Logger.Info("New task pushed")
}

func HandleDeleteTodo(c *gin.Context, storage *storage.Storage) {
	
}