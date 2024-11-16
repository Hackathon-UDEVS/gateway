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
 POSTGRES_PORT   int
 HTTP_PORT       string
 KAFKA_PORT      string
 MEDALS_HOST     string
 MEDALS_PORT     string
 KAFKA_HOST      string

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
  AUTHHOST:        cast.ToString(getEnv("AUTHHOST", "0")),
  GATEWAYHOST:     cast.ToString(getEnv("GATEWAYHOST", "0")),
  BOOKINGPORT:     cast.ToInt(getEnv("BOOKINGPORT", 1)),
  KAFKAPORT:       cast.ToInt(getEnv("KAFKAPORT", 1)),
  REDISPORT:       cast.ToInt(getEnv("REDISPORT", 1)),
  MONGOPORT:       cast.ToInt(getEnv("MONGOPORT", 1)),
  BOOKPORT:        cast.ToInt(getEnv("BOOKPORT", 1)),
  AUTHPORT:        cast.ToInt(getEnv("AUTHPORT", 1)),
  POSTGRES_PORT:   cast.ToInt(getEnv("POSTGRES_PORT", 1)),
  HTTP_PORT:       cast.ToString(getEnv("HTTP_PORT", 1)),
  KAFKA_HOST:      cast.ToString(getEnv("KAFKA_HOST", 1)),
  KAFKA_PORT:      cast.ToString(getEnv("KAFKA_PORT", 1)),
  MEDALS_HOST:     cast.ToString(getEnv("MEDALS_HOST", 1)),
  MEDALS_PORT:     cast.ToString(getEnv("MEDALS_PORT", 1)),
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
