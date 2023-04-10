package main

import (
	"net/http"

	"kantacky.com/api/handler/v1/domain"
	"kantacky.com/api/infrastructure"
)

func main() {
	ExecRouter()

	http.ListenAndServe("127.0.0.1:8080", nil)

	defer infrastructure.Close()
}

func ExecRouter() {
	http.HandleFunc("/api/v1/domain", domain.DomainHandler)
}
