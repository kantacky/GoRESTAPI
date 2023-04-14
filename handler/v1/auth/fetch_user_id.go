package auth

import (
	"fmt"
	"regexp"

	"kantacky.com/api/infrastructure"
)

func FetchUserId(username string, domain string) (string, error) {
	usernameRule := regexp.MustCompile(`[0-9a-z_\.]{1,32}`)
	domainRule := regexp.MustCompile(`[0-9a-z_\.]{1,128}`)
	if !usernameRule.MatchString(username) {
		return "", fmt.Errorf("Invalid username")
	}
	if !domainRule.MatchString(domain) {
		return "", fmt.Errorf("Invalid username")
	}

	query := fmt.Sprintf("SELECT id FROM identity.id WHERE username = '%s' AND domain = '%s' ORDER BY id ASC;", username, domain)
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
		return "", fmt.Errorf("User not found")
	}

	return idList[0], nil
}
