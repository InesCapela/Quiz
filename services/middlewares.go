package services

import (
	"Project_2021_PSRS/model"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)
func UserAuthorizationRequired() gin.HandlerFunc {
	return AuthorizationRequired(false)
}

func AdminAuthorizationRequired() gin.HandlerFunc {
	return AuthorizationRequired(true)
}
func AuthorizationRequired(requireAdmin bool) gin.HandlerFunc {

	return func(c *gin.Context) {
		if !ValidateTokenJWT(c) {

			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Not authorized"})
			c.Abort()
		} else {
			var tokenInput, _, _ = getAuthorizationToken(c)
			token, _ := jwt.ParseWithClaims(tokenInput, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
				return JwtKey, nil
			})

			if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
				//fmt.Printf("%v %v", claims.Username, claims.StandardClaims.ExpiresAt)
				c.Set("username", claims.Username)
				c.Set("isAdmin", claims.IsAdmin)
			}

			// Requires admin privileges
			if requireAdmin && !c.GetBool("isAdmin") {
				c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Not authorized"})
				c.Abort()
			}

			OpenDatabase()
			c.Next()

			defer Db.Close()
		}
	}
}
