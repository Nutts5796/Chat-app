package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Register(c *gin.Context, db *sqlx.DB) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	_, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User already exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

func Login(c *gin.Context, db *sqlx.DB) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var storedPassword string
	err := db.Get(&storedPassword, "SELECT password FROM users WHERE username=$1", user.Username)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)

	c.JSON(http.StatusOK, gin.H{"token": token})
}
