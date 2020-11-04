package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"

	"github.com/gorilla/mux"

	"github.com/siddhantk232/currency/protos/currency"
	"github.com/siddhantk232/go-micro/handlers"
)

func main() {
	l := log.New(os.Stdout, "[products-api] ", log.LstdFlags)

	currencyConnection, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	defer currencyConnection.Close()

	if err != nil {
		l.Println("error connecing to currency service", err)
	}

	cc := currency.NewCurrencyClient(currencyConnection)

	productsHandler := handlers.NewProducts(l, cc)

	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	postRouter := sm.Methods(http.MethodPost).Subrouter()

	putRouter.Use(productsHandler.ParseProductBody)
	postRouter.Use(productsHandler.ParseProductBody)

	getRouter.HandleFunc("/products", productsHandler.GetProducts)
	getRouter.HandleFunc("/products/{id:[0-9]+}", productsHandler.GetProduct)
	postRouter.HandleFunc("/products", productsHandler.AddProduct)
	putRouter.HandleFunc("/products/{id:[0-9]+}", productsHandler.UpdateProduct)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		l.Println("Server started. Listening on port 9090")
		error := server.ListenAndServe()
		if error != nil {
			l.Fatal(error)
		}
	}()

	tc, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	l.Println("\nShutting down the server", sig)
	server.Shutdown(tc)
	cancelFunc()
}
