package main

import (
	"errors"
	"fmt"
	"github.com/bayucandra/go-seed-pg/db"
	file_operations "github.com/bayucandra/go-seed-pg/file-operations"
	sql_operations "github.com/bayucandra/go-seed-pg/seed-operations"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	if os.Getenv("GO_ENV") == "development" || os.Getenv("GO_ENV") == "test" {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	envFile := ".env"

	if len(os.Args) > 1 {
		envFile = os.Args[1]
	}

	err := loadEnv( envFile)

	if err != nil {
		log.Println("Not loading .env var")
	}

	_, err = checkEnv()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(os.Getenv("GO_SEED_SOURCE_PATH"))

	db.Init()
	files, err := file_operations.DirParse(os.Getenv("GO_SEED_SOURCE_PATH"))
	sql_operations.SeedAll(files)

	_ = db.DBConn.Close()

}

func loadEnv(fileName string) error {
	err := godotenv.Load(fileName)
	return err
}

func checkEnv(envNames ...string) (notfound string, err error) {

	envVars := []string{
		"GO_SEED_PG_HOST",
		"GO_SEED_PG_PORT",
		"GO_SEED_PG_USER",
		"GO_SEED_PG_PASSWORD",
		"GO_SEED_PG_DBNAME",
		"GO_SEED_PG_SSLMODE",
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
