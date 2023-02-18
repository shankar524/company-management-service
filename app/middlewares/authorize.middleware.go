package middlewares

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"

	"github.com/gin-gonic/gin"
)

const BEARER_SCHEMA = "Bearer"

// Claims object
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Remove "Bearer " from "Authorization" token string
func stripBearer(tok string) (string, error) {
	if len(tok) > 6 && strings.ToLower(tok[0:7]) == "bearer " {
		return tok[7:], nil
	}
	return tok, nil
}

func Authorize(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := stripBearer(c.Request.Header.Get("Authorization"))
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenClaims, err := jwt.ParseWithClaims(
			token,
			&Claims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtSecret), nil
			},
		)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if tokenClaims != nil {
			claims, ok := tokenClaims.Claims.(*Claims)

			if ok && tokenClaims.Valid {
				// Set context values
				c.Set("username", claims.Username)

				c.Next()
				return
			}
		}

		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
