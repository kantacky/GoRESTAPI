package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"kantacky.com/api/handler/v1/auth"
	"kantacky.com/api/handler/v1/domain"
	"kantacky.com/api/handler/v1/user"
	"kantacky.com/api/handler/v1/util"
	"kantacky.com/api/infrastructure"
)

func main() {
	ExecRouter()

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	addr := fmt.Sprintf("%s:%s", host, port)
	log.Printf("Listening on %s", addr)

	http.ListenAndServe(addr, nil)

	defer infrastructure.Close()
}

func ExecRouter() {
	http.HandleFunc("/api/v1/util/uuid", util.UUIDHandler)
	http.HandleFunc("/api/v1/util/hash", util.HashHandler)

	http.HandleFunc("/api/v1/auth", auth.AuthHandler)

	http.HandleFunc("/api/v1/user", TokenRequired(user.UsersHandler))
	http.HandleFunc("/api/v1/user/", TokenRequired(user.UserHandler))
	http.HandleFunc("/api/v1/domain", domain.DomainHandler)
}

func TokenRequired(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		_, err := auth.VerifyToken(reqToken)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized: " + err.Error()))
			return
		}

		next(w, r)
	}
}
