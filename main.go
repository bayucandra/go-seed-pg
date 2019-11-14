package main

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := loadEnv(".env")

	if err != nil {
		log.Fatal(err)
	}

	_, err = checkEnv()

	if err != nil {
		log.Fatal(err)
	}

}

func loadEnv(fileName string) error {
	err := godotenv.Load(fileName)
	return err
}

func checkEnv(envNames ...string) (notfound string, err error) {

	envVars := []string{
		"PG_HOST",
		"PG_PORT",
		"PG_USER",
		"PG_PASSWORD",
		"PG_DBNAME",
		"PG_SSLMODE",
	}

	if len(envNames) > 0 {
		for _, val := range envNames {
			envVars = append(envVars, val)
		}
	}

	for _, val := range envVars {
		if os.Getenv(val) == "" {
			notfound = val
			err = errors.New(fmt.Sprintf("Env variable %s is not found", val))
			return
		}
	}

	return
}
