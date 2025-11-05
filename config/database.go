package config

import (
    "fmt"
    "log"
    "gin-quickstart/models" 
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "host=127.0.0.1 user=postgres password=123 dbname=gin_gorm port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    fmt.Println("Connected to PostgreSQL successfully!")

    // Auto-migrate tables
    db.AutoMigrate(&models.User{})

    DB = db
}
