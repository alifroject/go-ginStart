package models

import "time"

type User struct {
    ID          uint      `gorm:"primaryKey"`
    FirstName   string    `gorm:"not null"`
    LastName    string    `gorm:"not null"`
    Email       string    `gorm:"unique;not null"`
    Password    string    `gorm:"not null"`
    PhoneNumber *string
    Status      string    `gorm:"default:'active'"`
    Role        string    `gorm:"default:'user'"`
    CreatedAt   time.Time `gorm:"autoCreateTime"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
