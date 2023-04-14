package domain

import (
	"encoding/json"
	"net/http"
)

func DomainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		result, err := FetchDomain()
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
