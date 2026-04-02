package config

import (
	"fmt"
	"os"
)

type Config struct {
	PublicHost string
	Port       string

	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = NewConfig()

func NewConfig() *Config {
	return new(Config{
		PublicHost: LoadEnv("PUBLIC_HOST", "http://localhost"),
		Port:       LoadEnv("PORT", "8080"),
		DBUser:     LoadEnv("DB_USER", "root"),
		DBPassword: LoadEnv("DB_PASSWORD", "mypassword"),
		DBAddress:  fmt.Sprintf("%s:%s", LoadEnv("DB_HOST", "127.0.0.1"), LoadEnv("DB_PORT", "3306")),
		DBName:     LoadEnv("DB_NAME", "ushort"),
	})
}

func LoadEnv(variable string, fallback string) string {
	val, ok := os.LookupEnv(variable)
	if ok == true {
		return val
	}
	return fallback
}
