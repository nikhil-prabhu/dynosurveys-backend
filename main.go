package main

import (
        "log"
        "net/http"
        "os"
        "os/signal"
        "syscall"

        "github.com/joho/godotenv"
        "github.com/nikhil-prabhu/dynosurveys-backend/utils"
        "github.com/nikhil-prabhu/dynosurveys-backend/routes"
        "github.com/nikhil-prabhu/dynosurveys-backend/controllers"
)

func main() {
        // Handle server termination
        handleInterrupts()

        log.Println("Starting server...")
        // Close DB connections when server is killed
        defer utils.ClosePostgreClient(controllers.PDB)
        defer utils.CloseMongoClient(controllers.MDB, controllers.Ctx)

        // Load environment
        err := godotenv.Load()
        if err != nil {
                log.Fatalln(err)
        }
        // Port to run server on
        port := os.Getenv("PORT")

        // Handle routes
        http.Handle("/", routes.Handlers())

        // Serve
        log.Println("Server listening on port:", port)
        log.Fatal(http.ListenAndServe(":"+port, nil))
}

// handleInterrupts handles server termination
// in a somewhat graceful way. It handles the
// SIGINT and SIGTERM signals.
func handleInterrupts() {
        // Channel to notify about received signal
        ch := make(chan os.Signal)
        // Notify received signal
        signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

        // Goroutine that blocks till signal is
        // received on channel
        go func() {
                <- ch
                log.Println("Server terminated. Exiting...")
                os.Exit(0)
        }()
}
