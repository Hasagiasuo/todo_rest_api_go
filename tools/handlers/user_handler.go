package handlers

import (
	"net/http"
	"pet_pr/tools/storage"
	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context, storage *storage.Storage) {
	id := c.Param("id")
	user := storage.GetUserById(id)
	c.HTML(http.StatusOK, "user.tmpl", gin.H {
		"id": user.ID,
		"name": user.Name,
		"email": user.Email,
		"password": user.Password,
 	})
}

func PostUser(c *gin.Context) {

}

func GetUsers(c *gin.Context) {

}
