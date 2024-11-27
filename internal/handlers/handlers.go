package handlers

import (
	database "auth_go/internal/db"
	"auth_go/internal/models"
	"auth_go/internal/utils"
	"encoding/json"
	"net/http"
)

func GenerateTokensHandler(w http.ResponseWriter, r *http.Request) {
	var req models.TokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	accessToken, err := utils.GenerateJWT(req.UserID, req.IP)
	if err != nil {
		http.Error(w, "Не удалось сгенерировать аксесс токен", http.StatusInternalServerError)
		return
	}

	refreshToken := utils.GenerateRefreshToken()
	refreshHash, err := utils.HashToken(refreshToken)
	if err != nil {
		http.Error(w, "Не удалось хэшировать маркер обновления", http.StatusInternalServerError)
		return
	}

	err = database.SaveRefreshToken(req.UserID, refreshHash)
	if err != nil {
		http.Error(w, "Не удалось сохранить маркер обновлений", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func RefreshTokensHandler(w http.ResponseWriter, r *http.Request) {
	var req models.RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	storedHash, err := database.GetRefreshTokenHash(req.UserID)
	if err != nil || !utils.VerifyToken(req.RefreshToken, storedHash) {
		http.Error(w, "Неверный рефреш токен", http.StatusUnauthorized)
		return
	}

	if req.IP != req.CurrentIP {
		utils.SendWarningEmail(req.UserID)
	}

	accessToken, err := utils.GenerateJWT(req.UserID, req.CurrentIP)
	if err != nil {
		http.Error(w, "Не удалось сгенерировать аксесс токен", http.StatusInternalServerError)
		return
	}

	refreshToken := utils.GenerateRefreshToken()
	refreshHash, err := utils.HashToken(refreshToken)
	if err != nil {
		http.Error(w, "Не удалось хэшировать маркер обновления", http.StatusInternalServerError)
		return
	}

	err = database.SaveRefreshToken(req.UserID, refreshHash)
	if err != nil {
		http.Error(w, "Не удалось сохранить рефреш токен", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
