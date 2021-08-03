package util

import (
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func CreateUniqueID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

/**
 * .envファイルのローディング
 */
func EnvLoad() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
