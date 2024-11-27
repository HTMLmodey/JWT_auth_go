package main

import (
	"log"
	"net/http"

	database "auth_go/internal/db"
	"auth_go/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Не удалось инициализировать базу данных: %v", err)
	}
	defer database.DB.Close()

	r := mux.NewRouter()

	r.HandleFunc("/auth/tokens", handlers.GenerateTokensHandler).Methods("POST")
	r.HandleFunc("/auth/refresh", handlers.RefreshTokensHandler).Methods("POST")

	log.Println("Сервер работает на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
