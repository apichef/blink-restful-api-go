package middlewares

import "github.com/gin-gonic/gin"

func Authenticate() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"foo": "bar",
	})
}
