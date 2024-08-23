package utils

import (
	"log"
	"os"
)

// validate file .env
func CheckEnvVar(key string) {
	if os.Getenv(key) == "" {
		log.Fatalf("Variável de ambiente não definida: %s", key)
	}
}