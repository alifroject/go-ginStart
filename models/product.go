package models

import "time"

type Product struct {
    ID          uint      `gorm:"primaryKey"`
    Name        string    `gorm:"not null"`
    Category    string    `gorm:"not null"`
    Price       float64   `gorm:"not null"`
    Stock       int       `gorm:"default:0"`
    Description *string
    CreatedAt   time.Time `gorm:"autoCreateTime"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
