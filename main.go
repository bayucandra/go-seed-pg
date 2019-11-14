package main

import (
	"errors"
	"fmt"
	"github.com/bayucandra/go-seed/db"
	file_operations "github.com/bayucandra/go-seed/file-operations"
	sql_operations "github.com/bayucandra/go-seed/seed-operations"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := loadEnv(".env")

	if err != nil {
		log.Fatal(err)
	}

	_, err = checkEnv()

	if err != nil {
		log.Fatal(err)
	}

	db.Init()
	files, err := file_operations.DirParse(os.Getenv("SOURCE_PATH"))
	sql_operations.SeedAll(files)

	_ = db.DBConn.Close()

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
