package main

import (
	"SE_MIM22_WEBSHOP_LOGINSERVICE/handler"
	"log"
	"net/http"
	"time"
)

func main() {
	var serveMux = http.NewServeMux()
	serveMux.HandleFunc("/login", handler.Login)
	log.Printf("\n\n\tLOGINSERVICE\n\nAbout to listen on Port: 8441.\n\n" +
		"SUPPORTED REQUESTS:\n" +
		"GET:\n" +
		"Login on: http://127.0.0.1:8441/login requires a JSON Body with the following format: " +
		"{\"username\":\"test\",\"password\":\"test\"}")
	server := &http.Server{
		Addr:              ":8441",
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       3 * time.Second,
		Handler:           serveMux,
	}
	log.Fatal(server.ListenAndServe())
}
