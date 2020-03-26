package middleware

import (
	"aigram-backend/config"
	"errors"
	"net/http"
	"strings"

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
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		bearerToken := strings.Split(header, " ")
		if len(bearerToken) < 0 {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		tokenString := bearerToken[1]
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
