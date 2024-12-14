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
	c.JSON(http.StatusOK, "Success")
}

func HandlePushTodo(c *gin.Context, storage *storage.Storage) {
	type Req struct {
		UID string
		Title string
	}
	var answ Req
	if err := c.BindJSON(&answ); err != nil {
		storage.Logger.Info(fmt.Sprintf("Cannot parse answer: %s", err.Error()))
		c.JSON(http.StatusBadRequest, "Not update")
		return
	}
	if err := storage.PushTodoByUID(answ.Title, answ.UID); err != nil {
		storage.Logger.Info(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	storage.Logger.Info("New task pushed")
	c.JSON(http.StatusOK, "Pushed")
}

func HandleDeleteTodo(c *gin.Context, storage *storage.Storage) {
	type PJson struct {
		ID string
	}
	var res PJson
	if err := c.BindJSON(&res); err != nil {
		storage.Logger.Error(fmt.Sprintf("Cannot pars task by id: %s", err.Error()))
		c.JSON(http.StatusBadRequest, "Error delete")
		return
	}
	if err := storage.DeleteTodoByID(res.ID); err != nil {
		storage.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	storage.Logger.Info("Task deleted!");
	c.JSON(http.StatusOK, "Deleted")
}