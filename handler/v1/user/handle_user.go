package user

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	prefix := "/api/v1/user/"
	if strings.HasPrefix(r.URL.Path, prefix) {
		slug := r.URL.Path[len(prefix):]
		rule := regexp.MustCompile(`([0-9a-f]{8})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{12})`)
		if !rule.MatchString(slug) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid ID"))
			return
		}

		switch r.Method {
		case http.MethodGet:
			result, err := FetchUser(slug)
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
			id := uuid.New()
			res, _ := json.Marshal(id)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(res)
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
}
