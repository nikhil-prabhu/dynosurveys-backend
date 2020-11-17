package utils

import (
        "../models"
        "log"
        "os"
        "context"

        "github.com/jinzhu/gorm"
        _ "github.com/jinzhu/gorm/dialects/postgres"
        "github.com/joho/godotenv"
        "go.mongodb.org/mongo-driver/mongo"
        "go.mongodb.org/mongo-driver/mongo/options"
)

// NewPostgreClient connects to a PostgreSQL database
// and returns the connection object.
func NewPostgreClient() *gorm.DB {
        // Load PostgreSQL database URI from environment
        err := godotenv.Load()
        if err != nil {
                log.Fatalln(err)
        }

        postgreURI := os.Getenv("PostgreURI")

        // Connect to DB
        db, err := gorm.Open("postgres", postgreURI)
        if err != nil {
                log.Fatalln(err)
        }

        // Migrate the schema
        db.AutoMigrate(&models.User{})
        return db
}

// NewMongoClient connects to a MongoDB database
// and returns the connection object along with
// the context.
func NewMongoClient() (*mongo.Client, *context.Context) {
        // Load Mongo database URI from environment
        err := godotenv.Load()
        if err != nil {
                log.Fatalln(err)
        }

        mongoURI := os.Getenv("mongoURI")

        // Create client
        client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
        if err != nil {
                log.Fatalln(err)
        }

        // Get context
        ctx, _ := context.WithCancel(context.Background())
        err = client.Connect(ctx)
        if err != nil {
                log.Fatalln(err)
        }

        return client, &ctx
}
