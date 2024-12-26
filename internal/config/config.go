package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	RateLimitIP      int
	RateLimitToken   int
	BlockDuration    int
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Erro ao carregar o arquivo .env, utilizando valores padrão.")
	}

	rateLimitIP, _ := strconv.Atoi(getEnv("RATE_LIMIT_IP", "5"))
	rateLimitToken, _ := strconv.Atoi(getEnv("RATE_LIMIT_TOKEN", "10"))
	blockDuration, _ := strconv.Atoi(getEnv("BLOCK_DURATION", "300"))

	return &Config{
		RateLimitIP:    rateLimitIP,
		RateLimitToken: rateLimitToken,
		BlockDuration:  blockDuration,
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}