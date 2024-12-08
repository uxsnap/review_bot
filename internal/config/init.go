package config

import (
	"github.com/joho/godotenv"
)

func Init(path string) error {
	return godotenv.Load(path)
}
