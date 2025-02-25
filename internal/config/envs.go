package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                   string
	PublicHost             string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTSecret              string
	JWTExpirationInSeconds int64
	Addr                   string
	Pw                     string
	Db                     int64
	Enabled                bool
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		Port:                   getEnv("PORT", "8080"),
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		DBUser:                 getEnv("DB_USER", "root"),
		DBPassword:             getEnv("DB_PASSWORD", "mypassword"),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:                 getEnv("DB_NAME", "leskov"),
		JWTSecret:              getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
		Addr:    getEnv("REDIS_ADDR", "localhost:6379"),
		Pw:      getEnv("REDIS_PW", ""),
		Db:      getEnvAsInt("REDIS_DB", 0),
		Enabled: getEnvAsBool("REDIS_ENABLED", false),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseBool(value)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
