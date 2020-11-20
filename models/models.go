package models

import (
        jwt "github.com/dgrijalva/jwt-go"
        "github.com/jinzhu/gorm"
)

// User data structure
type User struct {
        gorm.Model

        Name            string
        Email           string `gorm:"type:varchar(100);unique_index"`
        Gender          string `json:"Gender"`
        Password        string `json:"Password"`
}

// Form structure
type Form struct {
        gorm.Model

        FormID          string `gorm:"type:varchar(100);primaryKey" json:"form_id"`
        UserID          uint `gorm:"foreignKey:UserID" json:"user_id"`
}

// Token structure
type Token struct {
        UserID  uint
        Name    string
        Email   string
        *jwt.StandardClaims
}

// Exception structure
type Exception struct {
        Message string `json:"message"`
}
