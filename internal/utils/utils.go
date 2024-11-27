package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key")

func GenerateJWT(userID, ip string) (string, error) {
	claims := &jwt.MapClaims{
		"user_id": userID,
		"ip":      ip,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(jwtKey)
}

func GenerateRefreshToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	var codded = base64.URLEncoding.EncodeToString(b)
	fmt.Println(codded)
	return codded
}

func HashToken(token string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	return string(hash), err
}

func VerifyToken(token, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(token)) == nil
}

func SendWarningEmail(userID string) {
	println("Предупреждение отправлено пользователю:", userID)
}
