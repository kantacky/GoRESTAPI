package user

import (
	"kantacky.com/api/infrastructure"
	"kantacky.com/api/model"
)

func FetchUsers() ([]model.UserInfo, error) {
	query := "SELECT id.id, id.username, id.domain, fullname.first_name, fullname.middle_name, fullname.last_name, email.email FROM identity.id LEFT JOIN identity.fullname ON fullname.id = id.id LEFT JOIN identity.email ON email.id = id.id ORDER BY id ASC;"
	rows, err := infrastructure.DbId.Query(query)
	if err != nil {
		return []model.UserInfo{}, err
	}
	defer rows.Close()

	var userList []model.UserInfo
	for rows.Next() {
		var user_info model.UserInfo

		if err := rows.Scan(
			&user_info.Id,
			&user_info.Username,
			&user_info.Domain,
			&user_info.FirstName,
			&user_info.MiddleName,
			&user_info.LastName,
			&user_info.Email,
		); err != nil {
			return []model.UserInfo{}, err
		}

		userList = append(userList, user_info)
	}

	return userList, nil
}
