package handlers

import (
	"net/http"
	"pet_pr/tools/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HPOST_login(c *gin.Context, storage *storage.Storage) {
	type UserInfo struct {
		Email 		string
		Password 	string
	}
	var user UserInfo
	c.BindJSON(&user)
	tmp := storage.GetUserByEmail(user.Email)
	if(user.Password == tmp.Password) {
		c.SetCookie("username", tmp.Name, 3600, "/", "localhost", false, true)
		c.SetCookie("uid", strconv.Itoa(tmp.ID), 3600, "/", "localhost", false, true)
		c.JSON(http.StatusAccepted, "Success")
	} else {
		c.JSON(http.StatusBadRequest, "Not login")
	}
}

func HGET_login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{})
}

func HPOST_signup(c *gin.Context, storage *storage.Storage) {
	type UserInfo struct {
		Username 	string
		Email 		string
		Password 	string
	}
	var user UserInfo
	c.BindJSON(&user)
	storage.CreateNewUser(user.Username, user.Email, user.Password)
	uid := storage.GetUserByName(user.Username)
	c.SetCookie("username", user.Username, 3600, "/", "localhost", false, true)
	c.SetCookie("uid", strconv.Itoa(uid.ID), 3600, "/", "localhost", false, true)
	c.JSON(http.StatusAccepted, "Success")
}

func HGET_signup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.tmpl", gin.H{})
}

func HPOST_logout(c *gin.Context) {
	c.SetCookie("username", "", 3600, "/", "localhost", false, true)
	c.SetCookie("uid", "", 3600, "/", "localhost", false, true)
}