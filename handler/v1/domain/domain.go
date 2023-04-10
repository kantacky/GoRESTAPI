package domain

import (
	"encoding/json"
	"net/http"

	"kantacky.com/api/infrastructure"
	"kantacky.com/api/model"
)

func DomainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		result, err := FetchDomain()

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

func FetchDomain() ([]model.Domain, error) {
	query := `SELECT domain
	FROM domain.domain
	WHERE display_order IS NOT NULL
	ORDER BY display_order ASC;`

	rows, err := infrastructure.DbId.Query(query)

	if err != nil {
		return []model.Domain{}, err
	}

	defer rows.Close()

	var users []model.Domain

	for rows.Next() {
		var (
			domain string
		)

		if err := rows.Scan(&domain); err != nil {
			return []model.Domain{}, err
		}

		user := model.Domain{
			Domain: domain,
		}

		users = append(users, user)
	}

	return users, nil
}
