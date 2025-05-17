// database/init.go
package database

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

func init() {
    if os.Getenv("ENV") != "production" {
        if err := godotenv.Load(); err != nil {
            log.Println("Warning: .env file not found, skipping...")
        }
    }
}
