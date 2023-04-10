package user

import (
	"encoding/json"
	"net/http"

	"kantacky.com/api/infrastructure"
	"kantacky.com/api/model"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		result, err := FetchUser()

		res, err := json.Marshal(result)

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

func FetchUser() ([]model.User, error) {
	db := infrastructure.GetDB("id")

	query := `SELECT id, username, domain
	FROM identity.id
	ORDER BY id ASC;`

	rows, err := db.Query(query)

	if err != nil {
		return []model.User{}, err
	}

	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var (
			id, username, domain string
		)

		if err := rows.Scan(&id, &username, &domain); err != nil {
			return []model.User{}, err
		}

		user := model.User{
			Id:       id,
			Username: username,
			Domain:   domain,
		}

		users = append(users, user)
	}

	return users, nil
}
