package controllers

import (
        "../models"
        "../utils"
        "encoding/json"
        "fmt"
        "net/http"
        "time"

        jwt "github.com/dgrijalva/jwt-go"
        "golang.org/x/crypto/bcrypt"
)

// ErrorResponse structure
type ErrorResponse struct {
        Err string
}

// error interface
type error interface {
        Error() string
}

// PostgreSQL DB Client
var pdb = utils.NewPostgreClient()

// Login attempts to log in a user and writes
// the response
func Login(w http.ResponseWriter, r *http.Request) {
        user := &models.User{}
        err := json.NewDecoder(r.Body).Decode(user)
        if err != nil {
                resp := map[string]interface{}{
                        "status": false,
                        "message": "Invalid request.",
                }
                return
        }
        resp := FindOne(user.Email, user.Password)
        json.NewEncoder(w).Encode(resp)
}

// FindOne searches the user database to check
// whether a user exists or not.
func FindOne(email, password string) map[string]interface{} {
        user := &models.User{}

        // Search the database
        if err := pdb.Where("Email = ?", email).First(user).Error; err != nil {
                resp := map[string]interface{}{
                        "status": false,
                        "message": "Email address not found.",
                }
                return resp
        }
        expiresAt := time.Now().Add(time.Minute * 100000).Unix()

        // Check password match
        err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
        if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
                resp := map[string]interface{}{
                        "status": false,
                        "message": "Incorrect password.",
                }
                return resp
        }

        // JWT token
        tk := &models.Token{
                UserID: user.ID,
                Name: user.Name,
                Email: user.Email,
                StandardClaims: &jwt.StandardClaims{
                        ExpiresAt: expiresAt,
                },
        }

        token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

        tokenString, err := token.SignedString([]byte("secret"))
        if err != nil {
                fmt.Println(err)
        }

        resp := map[string]interface{}{
                "status": false,
                "message": "Logged in.",
        }
        resp["token"] = tokenString
        resp["user"] = user

        return resp
}

// CreateUser creates a new user in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
        user := &models.User{}
        json.NewDecoder(r.Body).Decode(user)

        pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
        if err != nil {
                fmt.Println(err)
                err := ErrorResponse{
                        Err: "Password encryption failed.",
                }
                json.NewEncoder(w).Encode(err)
        }

        user.Password = string(pass)

        createdUser := pdb.Create(user)

        if createdUser.Error != nil {
                fmt.Println(createdUser.Error)
        }
        json.NewEncoder(w).Encode(createdUser)
}
