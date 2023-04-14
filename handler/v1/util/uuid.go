package util

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func UUIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s := r.URL.Query().Get("n")
		if !numCheck(s) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request: n must be a number"))
			return
		}

		n, err := strconv.Atoi(s)
		if err != nil {
			n = 1
		}

		if n < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request: n must be greater than 0"))
			return
		}

		uuidList := []uuid.UUID{}

		for i := 0; i < n; i++ {
			uuid := uuid.New()
			uuidList = append(uuidList, uuid)
		}

		res, err := json.Marshal(uuidList)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error: " + err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)

	case http.MethodPost:
		w.WriteHeader(http.StatusMethodNotAllowed)

	case http.MethodPut:
		w.WriteHeader(http.StatusMethodNotAllowed)

	case http.MethodDelete:
		w.WriteHeader(http.StatusMethodNotAllowed)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func numCheck(str string) bool {
	for _, r := range str {
		if '0' > r || r > '9' {
			return false
		}
	}
	return true
}
