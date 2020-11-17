package utils

import (
        "../models"
        "log"
        "os"

        "github.com/jinzhu/gorm"
        _ "github.com/jinzhu/gorm/dialects/postgres"
        "github.com/joho/godotenv"
)

// NewPostgreClient connects to a PostgreSQL database
// and returns the connection object.
func NewPostgreClient() *gorm.DB {
        // Load dbURI from environment
        err := godotenv.Load()
        if err != nil {
                log.Fatalln(err)
        }

        dbURI := os.Getenv("dbURI")

        // Connect to DB
        db, err := gorm.Open("postgres", dbURI)
        if err != nil {
                log.Fatalln(err)
        }

        // Migrate the schema
        db.AutoMigrate(&models.User{})
        return db
}
