package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Constants struct {
	PORT          string
	MONGO_URI     string
	DATABASE_NAME string
}

var Constant Constants

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Load .env failed")
	}
}

func init() {
	loadEnv() // Load .env file

	Constant.PORT = os.Getenv("PORT")
	Constant.MONGO_URI = os.Getenv("MONGO_URI")
	Constant.DATABASE_NAME = os.Getenv("DATABASE_NAME")
}
