package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("main: Start Testing==========")
	exitVal := m.Run()
	log.Println("main: End testing============")

	os.Exit(exitVal)
}

func TestLoadEnv(t *testing.T) {
	err := loadEnv(".env.example")
	assert.Nil(t, err, "Err should be nil")

	err = loadEnv(".env.not-exist")
	assert.NotNil(t, err, "Err should be not nil")
}

func TestCheckEnv(t *testing.T) {
	_, err := checkEnv()
	assert.Nil( t, err, "Should be nil" )

	_, err = checkEnv("ENV_NOT_EXIST")
	assert.NotNil( t, err, "Should be not nil" )

}
