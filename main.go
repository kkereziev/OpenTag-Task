package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"opentag/handlers"
	"opentag/helpers"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

// Port is passed as env variable during image build
var Port string

func main() {
	if Port == "" {
		Port = os.Args[2]
	}

	logger := log.New(os.Stdout, "opentag-task ", log.LstdFlags)
	codec := helpers.NewCodec()
	serveMux := mux.NewRouter()

	th := handlers.NewTranslationHandler(logger, codec)

	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	postRouter := serveMux.Methods(http.MethodPost).Subrouter()

	// GET
	getRouter.HandleFunc("/history", th.GetHistory)

	// POST
	postRouter.HandleFunc("/word", th.TranslateWord)
	postRouter.HandleFunc("/sentence", th.TranslateSentence)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", Port),
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}
