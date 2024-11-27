package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	database "auth_go/internal/db"
	"auth_go/internal/handlers"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/auth/tokens", handlers.GenerateTokensHandler).Methods("POST")
	r.HandleFunc("/auth/refresh", handlers.RefreshTokensHandler).Methods("POST")
	return r
}

func TestGenerateTokensEndpoint(t *testing.T) {
	err := database.InitDB()
	assert.NoError(t, err, "База данных должна инициализироваться без ошибок")
	defer database.DB.Close()

	router := setupRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	reqBody := map[string]string{
		"user_id": "550e8400-e29b-41d4-a716-446655440000",
		"ip":      "192.168.1.10",
	}
	body, _ := json.Marshal(reqBody)

	resp, err := http.Post(server.URL+"/auth/tokens", "application/json", bytes.NewBuffer(body))
	assert.NoError(t, err, "Запрос не должен возвращать ошибку")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Статус кода должен быть 200")

	var respBody map[string]string
	json.NewDecoder(resp.Body).Decode(&respBody)
	assert.NotEmpty(t, respBody["access_token"], "Токен доступа не должен быть пустым")
	assert.NotEmpty(t, respBody["refresh_token"], "Рефреш токен не должен быть пустым")
}

func TestRefreshTokensEndpoint(t *testing.T) {
	err := database.InitDB()
	assert.NoError(t, err, "База данных должна инициализироваться без ошибок")
	defer database.DB.Close()

	router := setupRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	reqBody := map[string]string{
		"user_id": "550e8400-e29b-41d4-a716-446655440000",
		"ip":      "192.168.1.10",
	}
	body, _ := json.Marshal(reqBody)

	resp, err := http.Post(server.URL+"/auth/tokens", "application/json", bytes.NewBuffer(body))
	assert.NoError(t, err, "Запрос не должен возвращать ошибку")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Статус кода должен быть 200")

	var respBody map[string]string
	json.NewDecoder(resp.Body).Decode(&respBody)
	refreshToken := respBody["refresh_token"]

	refreshReqBody := map[string]string{
		"user_id":       "550e8400-e29b-41d4-a716-446655440000",
		"refresh_token": refreshToken,
		"ip":            "192.168.1.10",
		"current_ip":    "192.168.1.10",
	}
	refreshBody, _ := json.Marshal(refreshReqBody)

	refreshResp, err := http.Post(server.URL+"/auth/refresh", "application/json", bytes.NewBuffer(refreshBody))
	assert.NoError(t, err, "Запрос не должен возвращать ошибку")
	assert.Equal(t, http.StatusOK, refreshResp.StatusCode, "Статус кода должен быть 200")

	var refreshRespBody map[string]string
	json.NewDecoder(refreshResp.Body).Decode(&refreshRespBody)
	assert.NotEmpty(t, refreshRespBody["access_token"], "Новый токен не должен быть пустым")
	assert.NotEmpty(t, refreshRespBody["refresh_token"], "Новый токен не должен быть пустым")
}
