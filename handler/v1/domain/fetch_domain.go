package domain

import (
	"kantacky.com/api/infrastructure"
)

func FetchDomain() ([]string, error) {
	query := `SELECT domain
	FROM domain.domain
	WHERE display_order IS NOT NULL
	ORDER BY display_order ASC;`

	rows, err := infrastructure.DbId.Query(query)
	if err != nil {
		return []string{}, err
	}

	defer rows.Close()

	var domains []string

	for rows.Next() {
		var domain string

		if err := rows.Scan(&domain); err != nil {
			return []string{}, err
		}

		domains = append(domains, domain)
	}

	return domains, nil
}
