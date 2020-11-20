package auth

import (
        "context"
        "encoding/json"
        "net/http"
        "strings"

        jwt "github.com/dgrijalva/jwt-go"
        "github.com/nikhil-prabhu/dynosurveys-backend/models"
)

// Context key
type key string

// JWTVerify verifies the JWT token and returns
// an HTTP handler object.
func JWTVerify(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                header := r.Header.Get("x-access-token") // Get the token from the response header
                header = strings.TrimSpace(header)

                if header == "" {
                        // Token is missing. Return error code 403 unauthorized
                        w.WriteHeader(http.StatusForbidden)
                        json.NewEncoder(w).Encode(models.Exception{
                                Message: "Missing auth token.",
                        })
                        return
                }
                tk := &models.Token{}

                _, err := jwt.ParseWithClaims(header, tk,
                        func(token *jwt.Token) (interface{}, error) {
                                return []byte("secret"), nil
                        })
                if err != nil {
                        w.WriteHeader(http.StatusForbidden)
                        json.NewEncoder(w).Encode(models.Exception{
                                Message: err.Error(),
                        })
                        return
                }

                var ctxKey key = "user"
                ctx := context.WithValue(r.Context(), ctxKey, tk)
                next.ServeHTTP(w, r.WithContext(ctx))
        })
}
