package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	userID := "550e8400-e29b-41d4-a716-446655440000"
	ip := "192.168.1.10"

	token, err := GenerateJWT(userID, ip)
	assert.NoError(t, err, "Ошибка должна быть нулевой")
	assert.NotEmpty(t, token, "Токен не должен быть пустым")
}

func TestVerifyToken(t *testing.T) {
	token := GenerateRefreshToken()
	hash, err := HashToken(token)
	assert.NoError(t, err, "Хеширование не должно возвращать ошибку")

	isValid := VerifyToken(token, hash)
	assert.True(t, isValid, "Токен должен быть действительным")

	isInvalid := VerifyToken("wrong_token", hash)
	assert.False(t, isInvalid, "Токен должен быть недействительным")
}
