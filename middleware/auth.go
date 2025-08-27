package middleware

import (
    "net/http"
    "strings"

    "go-restapi-gin/models"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("kode_rahasia")

type Claims struct {
    ID       uint   `json:"id"`
    Username string `json:"username"`
    jwt.RegisteredClaims
}

func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        var tokenString string

        // cek header Authorization
        header := c.GetHeader("Authorization")
        if header != "" && strings.HasPrefix(header, "Bearer ") {
            tokenString = header[7:]
        } else {
            // fallback cookie
            cookie, err := c.Cookie("token")
            if err == nil {
                tokenString = cookie
            }
        }

        if tokenString == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token tidak ditemukan"})
            return
        }

        claims := &Claims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })
        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token tidak valid"})
            return
        }

        // ambil user dari DB (opsional — untuk memastikan token sama dengan yang tersimpan)
        var user models.User
        if err := models.DB.First(&user, claims.ID).Error; err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user tidak ditemukan"})
            return
        }

        // optional: cek token di DB sama
        if user.Token != tokenString {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token tidak cocok"})
            return
        }

        // simpan user di context
        c.Set("user", user)
        c.Next()
    }
}
