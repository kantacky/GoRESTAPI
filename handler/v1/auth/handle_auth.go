package auth

import (
	"encoding/json"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		username := r.Header.Get("Username")
		domain := r.Header.Get("Domain")
		password := r.Header.Get("Password")

		result, err := AuthUser(username, domain, password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error: " + err.Error()))
			return
		}

		res, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error: " + err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		break

	case http.MethodPost:
		w.WriteHeader(http.StatusMethodNotAllowed)
		break

	case http.MethodPut:
		w.WriteHeader(http.StatusMethodNotAllowed)
		break

	case http.MethodDelete:
		w.WriteHeader(http.StatusMethodNotAllowed)
		break

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		break
	}
}
