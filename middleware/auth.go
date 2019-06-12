package middleware

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)


func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		username := session.Get("user")

		if username == "" || username == nil {
			c.Redirect(302, "/login")
			return
		}

		c.Next()
	}
}
