package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("C25.j7e3")

type Claims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// GenerarJWT genera un nuevo token para un usuario autenticado
func GenerarJWT(id uint, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ID:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// AuthMiddleware verifica el token JWT en las solicitudes protegidas
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Falta el token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Set("profesorID", claims.ID)
		c.Next()
	}
}
