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

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Listening on %s", addr)

	http.ListenAndServe(addr, nil)

	defer infrastructure.Close()
}

func ExecRouter() {
	http.HandleFunc("/api/v1/domain", domain.DomainHandler)
}
