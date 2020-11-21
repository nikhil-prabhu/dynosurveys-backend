package main

import (
        "log"
        "net/http"
        "os"
        "os/signal"

        "github.com/joho/godotenv"
        "github.com/nikhil-prabhu/dynosurveys-backend/controllers"
        "github.com/nikhil-prabhu/dynosurveys-backend/routes"
        "github.com/nikhil-prabhu/dynosurveys-backend/utils"
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
// in a somewhat graceful way.
func handleInterrupts() {
        // Channel to notify about received signal.
        // We use a buffered channel to make sure
        // that it's ready to receive when a signal
        // is sent
        ch := make(chan os.Signal, 1)
        // Notify received signal
        signal.Notify(ch)

        // Goroutine that blocks till signal is
        // received on channel
        go func() {
                <-ch
                log.Println("Server terminated. Exiting...")
                // Ensure that DB connections are closed before exiting
                utils.ClosePostgreClient(controllers.PDB)
                utils.CloseMongoClient(controllers.MDB, controllers.Ctx)
                os.Exit(0)
        }()
}
