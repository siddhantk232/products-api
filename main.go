package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/siddhantk232/go-micro/handlers"
)

func main() {
	l := log.New(os.Stdout, "products-api ", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	productsHandler := handlers.NewProducts(l)

	sm := http.NewServeMux()

	sm.Handle("/", helloHandler)
	sm.Handle("/products", productsHandler)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		error := server.ListenAndServe()
		if error != nil {
			l.Fatal(error)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	l.Println("\nShutting down the server", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	server.Shutdown(tc)

}
