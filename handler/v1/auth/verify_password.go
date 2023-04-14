package auth

import (
	"fmt"

	"kantacky.com/api/infrastructure"
)

func VerifyPassword(id string, passwordHash string) (string, error) {
	query := fmt.Sprintf("SELECT id FROM identity.password WHERE id = '%s' AND password_hash = '%s' ORDER BY id ASC;", id, passwordHash)
	rows, err := infrastructure.DbId.Query(query)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var idList []string

	for rows.Next() {
		var id string

		if err := rows.Scan(&id); err != nil {
			return "", err
		}

		idList = append(idList, id)
	}

	if len(idList) != 1 {
		return "", fmt.Errorf("Password is incorrect")
	}

	return idList[0], nil
}
