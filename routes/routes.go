package routes

import (
        "../controllers"
        "net/http"

        "github.com/gorilla/mux"
)

// Handlers defines and returns a router object
func Handlers() *mux.Router {
        r := mux.NewRouter().StrictSlash(true)
        r.Use(CommonMiddleware)

        r.HandleFunc("/register", controllers.CreateUser).Methods("POST")
        r.HandleFunc("/login", controllers.Login).Methods("POST")
        r.HandleFunc("/record_response", controllers.RecordFormResponse).Methods("POST")

        return r
}

// CommonMiddleware defines and returns a
// middleware handler object.
func CommonMiddleware(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.Header().Add("Content-Type", "application/json")
                w.Header().Set("Access-Control-Allow-Origin", "*")
                w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
                w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
                next.ServeHTTP(w, r)
        })
}