package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func inArray(arr []string, value string) bool {
	inarr := false

	for _, v := range arr {
		if v == value {
			inarr = true
			break
		}
	}

	return inarr
}

// CSRF protection
func CSRF() gin.HandlerFunc {
	ignoreMethods := []string{"GET", "HEAD", "OPTIONS"}
	return func(c *gin.Context) {
		if inArray(ignoreMethods, c.Request.Method) {
			c.Next()
			return
		}
		headerToken := c.GetHeader("X-CSRF-TOKEN")
		cookieToken, err := c.Cookie("CSRF-TOKEN")
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		if headerToken != cookieToken {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
}
