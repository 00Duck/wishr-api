package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/00Duck/wishr-api/app"
)

func main() {
	env := app.New()
	env.DB.Connect()

	port := "9191"
	svr := &http.Server{
		Addr:         "127.0.0.1:" + port,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Second * 120,
		Handler:      env.Router,
	}

	go func() {
		log.Print("Starting API on port " + port)
		err := svr.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// block until we receive a kill or interrupt
	<-c

	//when the signal is caught, gracefully shutdown the server
	//Allows a timeout for processes to complete if any are in progress
	_, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	//db.Disconnect(ctx)
}
