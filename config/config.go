package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	PublicHost string
	Port       string

	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

var Envs = NewConfig()

func NewConfig() *Config {
	return new(Config{
		PublicHost:             LoadEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   LoadEnv("PORT", "8080"),
		DBUser:                 LoadEnv("DB_USER", "root"),
		DBPassword:             LoadEnv("DB_PASSWORD", "mypassword"),
		DBAddress:              fmt.Sprintf("%s:%s", LoadEnv("DB_HOST", "127.0.0.1"), LoadEnv("DB_PORT", "3306")),
		DBName:                 LoadEnv("DB_NAME", "ushort"),
		JWTSecret:              LoadEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
		JWTExpirationInSeconds: LoadIntEnv("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
	})
}

func LoadEnv(variable string, fallback string) string {
	val, ok := os.LookupEnv(variable)
	if ok == true {
		return val
	}
	return fallback
}

func LoadIntEnv(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}

	return fallback
}
