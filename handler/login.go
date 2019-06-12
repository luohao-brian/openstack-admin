package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
)

func LoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", gin.H{})
	return
}

func LoginPostHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username != "admin" || password != "admin" {
		c.Redirect(302, "/login")	
		return
	}

	session := sessions.Default(c)
	session.Set("user", username)	
	session.Save()

	c.Redirect(302, "/")	
	return
}
