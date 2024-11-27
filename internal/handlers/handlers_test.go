package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	database "auth_go/internal/db"
	"auth_go/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTokensHandler(t *testing.T) {
	err := database.InitDB()
	assert.NoError(t, err, "База данных должна инициализироваться без ошибок")
	defer database.DB.Close()

	reqBody := models.TokenRequest{
		UserID: "550e8400-e29b-41d4-a716-446655440000",
		IP:     "192.168.1.10",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("POST", "/auth/tokens", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	GenerateTokensHandler(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Статус код должен быть 200")

	var respBody models.TokenResponse
	json.NewDecoder(resp.Body).Decode(&respBody)

	assert.NotEmpty(t, respBody.AccessToken, "Аксесс токен не должен быть пустым")
	assert.NotEmpty(t, respBody.RefreshToken, "Рефреш токен не должен быть пустым")
}
