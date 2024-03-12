package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_DSN                 string
	DBUser                 string
	DBPassword             string
	DBName                 string
	InstanceConnectionName string
	Mode                   string
)

// .envを呼び出します。
func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("読み込み出来ませんでした: %v", err)
	}

	DB_DSN = os.Getenv("DB_DSN")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	InstanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME")
	Mode = os.Getenv("MODE")
}
