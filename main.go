package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/00Duck/wishr-api/app"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	env := app.New()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:80", "https://localhost:443", "https://localhost:8080", "http://localhost:8080"},
		AllowCredentials: true,
	})

	port := "9191"
	svr := &http.Server{
		Addr:         "127.0.0.1:" + port,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Second * 120,
		Handler:      c.Handler(env.Router),
	}

	go func() {
		env.Log.Printf("Starting API on port " + port)
		err := svr.ListenAndServe()
		if err != nil {
			env.Log.Fatal(err)
		}
	}()

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	signal.Notify(channel, os.Kill)

	// block until we receive a kill or interrupt
	<-channel

	//when the signal is caught, gracefully shutdown the server
	//Allows a timeout for processes to complete if any are in progress
	_, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
}
