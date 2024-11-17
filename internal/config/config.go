package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config holds the configuration values.
type Config struct {
	DBHOST          string
	DBPORT          int
	DBUSER          string
	DBPASSWORD      string
	DBNAME          string
	KAFKAHOST       string
	KAFKAPORT       int
	REDISHOST       string
	REDISPORT       int
	TOKENKEY        string
	EMAILSECREDKEY  string
	EMAIL           string
	MONGOPORT       int
	MONGOHOST       string
	MONGODBDATABASE string
	BOOKINGHOST     string
	BOOKINGPORT     int
	BOOKHOST        string
	BOOKPORT        int
	AUTHPORT        int
	AUTHHOST        string
	GATEWAYHOST     string
	GATEWAYPORT     int
	TENDERHOST      string
	TENDERPORT      int
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)

	}

	// Populate configuration struct
	config := Config{
		DBHOST:          cast.ToString(getEnv("POSTGRESHOST", "1")),
		DBPORT:          cast.ToInt(getEnv("POSTGRESPORT", 1)),
		DBUSER:          cast.ToString(getEnv("POSTGRESUSER", "1")),
		DBPASSWORD:      cast.ToString(getEnv("POSTGRESPASSWORD", "1")),
		DBNAME:          cast.ToString(getEnv("POSTGESDB", "1")),
		REDISHOST:       cast.ToString(getEnv("REDISHOST", "0")),
		EMAILSECREDKEY:  cast.ToString(getEnv("EMAILSECREDKEY", "0")),
		EMAIL:           cast.ToString(getEnv("EMAIL", "0")),
		MONGOHOST:       cast.ToString(getEnv("MONGOHOST", "0")),
		MONGODBDATABASE: cast.ToString(getEnv("MONGODBDATABASE", "0")),
		BOOKINGHOST:     cast.ToString(getEnv("BOOKINGHOST", "0")),
		BOOKHOST:        cast.ToString(getEnv("BOOKHOST", "0")),
		AUTHHOST:        cast.ToString(getEnv("AUTHHOST", "localhost")),
		BOOKINGPORT:     cast.ToInt(getEnv("BOOKINGPORT", 1)),
		TOKENKEY:        cast.ToString(getEnv("TOKENKEY", "this_is_secred")),
		GATEWAYHOST:     cast.ToString(getEnv("GATEWAYHOST", "localhost")),
		KAFKAPORT:       cast.ToInt(getEnv("KAFKAPORT", 1)),
		REDISPORT:       cast.ToInt(getEnv("REDISPORT", 1)),
		MONGOPORT:       cast.ToInt(getEnv("MONGOPORT", 1)),
		BOOKPORT:        cast.ToInt(getEnv("BOOKPORT", 1)),
		AUTHPORT:        cast.ToInt(getEnv("AUTHPORT", 8085)),
		GATEWAYPORT:     cast.ToInt(getEnv("GATEWAYPORT", 8091)),
		TENDERHOST:      cast.ToString(getEnv("AUTHPORT", "localhost")),
		TENDERPORT:      cast.ToInt(getEnv("GATEWAYPORT", 8080)),
	}

	return config
}

// getEnv retrieves an environment variable or returns a default value.
func getEnv(key string, defaultVal interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if !exists {
		return defaultVal
	}
	return val
}
