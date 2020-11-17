package models

import (
        "github.com/jinzhu/gorm"
        jwt "github.com/dgrijalva/jwt-go"
)

// User data structure
type User struct {
        gorm.Model

        Name            string
        Email           string `gorm:"type:varchar(100);unique_index"`
        Gender          string `json:"Gender"`
        Password        string `json:"Password"`
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
