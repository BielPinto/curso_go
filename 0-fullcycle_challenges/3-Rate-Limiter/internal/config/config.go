package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisAddr     string
	RedisPassword string
	RedisDB       int

	RateLimitIP    int
	RateLimitToken int
	BlockTimeIP    int // in seconds
	BlockTimeToken int // in seconds

	Port string
}

func Load() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := &Config{
		RedisAddr:      getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:  getEnv("REDIS_PASSWORD", ""),
		RedisDB:        getEnvAsInt("REDIS_DB", 0),
		RateLimitIP:    getEnvAsInt("RATE_LIMIT_IP", 10),
		RateLimitToken: getEnvAsInt("RATE_LIMIT_TOKEN", 100),
		BlockTimeIP:    getEnvAsInt("BLOCK_TIME_IP", 300),    // 5 minutes
		BlockTimeToken: getEnvAsInt("BLOCK_TIME_TOKEN", 300), // 5 minutes
		Port:           getEnv("PORT", "8080"),
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
