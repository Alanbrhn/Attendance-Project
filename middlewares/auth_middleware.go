package middlewares

import (
	"Attendance-Project/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mengambil token dari header Authorization
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		// Menghapus prefix "Bearer"
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Verifikasi token menggunakan fungsi VerifyJWT
		_, err := utils.VerifyJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Lanjutkan ke request berikutnya
		c.Next()
	}
}
