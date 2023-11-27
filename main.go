package main

import (
	"context"
	"golangApp_docker_ks8/handlers"
	"golangApp_docker_ks8/version"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Printf(
		"Starting the service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)

	port := os.Getenv("PORT")
	log.Printf("Port is %s\n", port)
	if port == "" {
		log.Fatal("Port is not set.")
	}

	router := handlers.Route(version.BuildTime, version.Commit, version.Release)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	log.Print("The services is ready to listen and serve")

	killSignal := <-interrupt
	switch killSignal {
	case os.Interrupt:
		log.Print("Got SIGNT ...")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM ...")
	}

	log.Print("Service is shutting down ...")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Print("Error shutdown: ", err)
	}
	log.Print("Done")
}
