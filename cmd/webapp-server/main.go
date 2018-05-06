package main

import (
	"github.com/syunkitada/go-gin-sample/handler"
	"log"
	"net/http"
	"time"
)

func main() {
	app_handler := handler.GetHandler()

	s := &http.Server{
		Addr:           ":8000",
		Handler:        app_handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
