package env

import (
	"os"

	"github.com/joho/godotenv"
)

func GetOrDie(key string) string {
	if value, success := os.LookupEnv(key); success {
		return value
	}
	panic("Environment variable " + key + " is not set")
}

func LoadEnv() {
	godotenv.Load("prv.env")
}
