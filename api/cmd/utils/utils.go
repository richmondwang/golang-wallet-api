package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/richmondwang/golang-wallet-api/ent"
)

func GetDBClient() (*ent.Client, error) {
	dbDataSourceSlice := []interface{}{
		GetEnvOrFail("DB_HOST"),
		GetEnvOrFail("DB_USER"),
		GetEnvOrFail("DB_PASSWORD"),
		GetEnvOrFail("DB_NAME"),
		GetEnvOrFail("DB_PORT"),
	}
	return ent.Open("postgres", fmt.Sprintf("host=%[1]s user=%[2]s password=%[3]s dbname=%[4]s port=%[5]s sslmode=disable", dbDataSourceSlice...))
}

func GetEnvOrFail(envName string) string {
	val, x := os.LookupEnv(envName)
	if !x {
		log.Fatalf("ENV VAR %s is required", envName)
	}
	return val
}

func GetEnvOrDefault(envName, defaultVal string) string {
	val, x := os.LookupEnv(envName)
	if !x {
		return defaultVal
	}
	return val
}
