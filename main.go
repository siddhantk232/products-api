package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/siddhantk232/products-api/setup"
)

func main() {
	logger := log.New(os.Stdout, "products-api ", log.LstdFlags)

	sm := mux.NewRouter()

	setup.SetupRoutes(sm)

	server := http.Server{
		Addr:    ":9090",
		Handler: sm,
	}

	go func() {
		logger.Println("Server started on port :9090")
		err := server.ListenAndServe()

		if err != nil {
			logger.Printf("Error starting the server %s", err.Error())
			os.Exit(1)
		}
	}()

	// graceful shutdown setup
	tc, cancelFunc := context.WithTimeout(context.Background(), time.Second*60)
	defer cancelFunc()

	sigChan := make(chan os.Signal)

	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	logger.Println("\n Shutting down the server", sig)
	server.Shutdown(tc)
}
