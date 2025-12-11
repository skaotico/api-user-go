package security

import (
	jwtPlatform "api-user-go/pkg/platform/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	jwtGo "github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware crea un middleware para validar  tokens JWT.
func AuthMiddleware(jwtManager *jwtPlatform.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "El encabezado de autorización es obligatorio"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de encabezado de autorización inválido"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwtManager.Validate(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwtGo.MapClaims)
		if ok {
			if userID, ok := claims["user_id"]; ok {
				c.Set("user_id", userID)
			}
			if email, ok := claims["email"]; ok {
				c.Set("email", email)
			}
		}

		c.Next()
	}
}
