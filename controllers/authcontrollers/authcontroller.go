package authcontrollers

import (
    "net/http"
    "time"

    "go-restapi-gin/models"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("kode_rahasia") // gunakan env var di production
const tokenExpireSeconds = 3600 * 24 // 24 jam

type Claims struct {
    ID       uint   `json:"id"`
    Username string `json:"username"`
    jwt.RegisteredClaims
}

// Register handler
func Register(c *gin.Context) {
    var input struct {
        Username string `json:"username" binding:"required"`
        Password string `json:"password" binding:"required"`
        Email    string `json:"email" binding:"required,email"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // hash password
    hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal hash password"})
        return
    }

    user := models.User{
        Username: input.Username,
        Password: string(hashed),
        Email:    input.Email,
    }

    if err := models.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "register berhasil", "user": gin.H{"id": user.ID, "username": user.Username, "email": user.Email}})
}

// Login handler -> buat token, simpan ke DB dan set cookie
func Login(c *gin.Context) {
    var input struct {
        Username string `json:"username" binding:"required"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := models.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "username salah"})
        return
    }

    // compare password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "password salah"})
        return
    }

    // buat token JWT
    expirationTime := time.Now().Add(time.Second * tokenExpireSeconds)
    claims := &Claims{
        ID:       user.ID,
        Username: user.Username,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            Subject:   "auth",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal membuat token"})
        return
    }

    // simpan token di DB
    user.Token = tokenString
    models.DB.Save(&user)

    // set cookie (expired in tokenExpireSeconds)
    // domain "/" dan secure false untuk development; ganti secure=true & domain sesuai production
    c.SetCookie("token", tokenString, tokenExpireSeconds, "/", "localhost", false, true)

    c.JSON(http.StatusOK, gin.H{"message": "login berhasil", "token": tokenString})
}

// Logout -> hapus token di DB & cookie
func Logout(c *gin.Context) {
    // ambil user dari context jika middleware sudah memasukkan
    u, exists := c.Get("user")
    if exists {
        user := u.(models.User)
        user.Token = ""
        models.DB.Save(&user)
    }
    // hapus cookie dengan memberi maxAge negatif
    c.SetCookie("token", "", -1, "/", "localhost", false, true)
    c.JSON(http.StatusOK, gin.H{"message": "logout berhasil"})
}
