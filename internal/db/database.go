package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("postgres", "user=user password=password dbname=postgres sslmode=disable")
	if err != nil {
		return err
	}
	return DB.Ping()
}

func SaveRefreshToken(userID, hash string) error {
	_, err := DB.Exec("INSERT INTO tokens (user_id, refresh_hash) VALUES ($1, $2) ON CONFLICT (user_id) DO UPDATE SET refresh_hash = $2", userID, hash)
	return err
}

func GetRefreshTokenHash(userID string) (string, error) {
	var hash string
	err := DB.QueryRow("SELECT refresh_hash FROM tokens WHERE user_id = $1", userID).Scan(&hash)
	return hash, err
}
