package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveAndGetRefreshToken(t *testing.T) {
	err := InitDB()
	assert.NoError(t, err, "База данных должна инициализироваться без ошибок")
	defer DB.Close()

	userID := "550e8400-e29b-41d4-a716-446655440000"
	hash := "test_hash"

	err = SaveRefreshToken(userID, hash)
	assert.NoError(t, err, "Сохранение маркера обновления не должно возвращать ошибку")

	storedHash, err := GetRefreshTokenHash(userID)
	assert.NoError(t, err, "Получение рефреш токена не должно возвращать ошибку")
	assert.Equal(t, hash, storedHash, "Сохраненный хэш должен совпадать с оригиналом")
}
