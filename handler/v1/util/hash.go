package util

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

func HashHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var raw_string_list []string

		err := json.NewDecoder(r.Body).Decode(&raw_string_list)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		hashList := []string{}

		for _, raw_string := range raw_string_list {
			hashList = append(hashList, GetHashValueOf(raw_string))
		}

		res, err := json.Marshal(hashList)

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

func GetHashValueOf(raw_string string) string {
	hash := sha512.Sum512([]byte(raw_string))
	hash_value := hex.EncodeToString(hash[:])
	return hash_value
}
