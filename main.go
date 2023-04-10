package main

import (
	"fmt"
	"log"
	"net/http"

	"kantacky.com/api/handler/v1/domain"
	"kantacky.com/api/infrastructure"
)

func main() {
	ExecRouter()

	host := "127.0.0.1"
	port := 8080
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("Listening on http://%s", addr)

	http.ListenAndServe(addr, nil)

	defer infrastructure.Close()
}

func ExecRouter() {
	http.HandleFunc("/api/v1/domain", domain.DomainHandler)
}
