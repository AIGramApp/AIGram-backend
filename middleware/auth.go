package middleware

import (
	"aigram-backend/config"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Claims is used to parse jwt information
type Claims struct {
	ID int64 `json:"id"`
	jwt.StandardClaims
}

// AuthenticationRequired is used to verify the user token is present on the request
// If it is present, jwt is parsed and user information is set on the context
func AuthenticationRequired(config *config.AppConfiguration) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie(config.JWT.CookieName)
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		claims := Claims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(tkn *jwt.Token) (interface{}, error) {
			if tkn.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, errors.New("Wrong signging algorithm")
			}
			return []byte(config.JWT.Secret), nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		if !token.Valid {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Set("currentUser", claims)
		c.Next()
	}
}
