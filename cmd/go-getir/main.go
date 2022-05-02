package main

import (
	"context"
	"go-getir/internal/setup"
	"go-getir/pkg/config"
	"go-getir/pkg/mongo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func initServer(cfg *config.Configuration) *http.Server {

	mc := mongo.NewMongoClient(cfg.MongoDB.Uri)

	services := setup.SetupServices(cfg, mc)

	routes := setup.SetupRoutes(services)

	server := &http.Server{
		Addr:         ":" + cfg.Server.Port, // configure the bind address
		ReadTimeout:  50 * time.Second,      // max time to read request from the client
		WriteTimeout: 100 * time.Second,     // max time to write response to the client
		IdleTimeout:  12 * time.Second,      // max time for connections using TCP Keep-Alive
	}

	server.Handler = routes

	return server
}

func main() {
	if err := config.Setup(); err != nil {
		log.Fatalf("config.Setup() error: %s", err)
	}
	cfg := config.GetConfig()

	server := initServer(cfg)

	// start the server
	go func() {
		log.Println("Starting server on", server.Addr)

		err := server.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 10 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_ = server.Shutdown(ctx)
}
