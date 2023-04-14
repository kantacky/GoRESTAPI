package user

import (
	"fmt"

	"kantacky.com/api/infrastructure"
	"kantacky.com/api/model"
)

func FetchUser(user_id string) (model.UserInfo, error) {
	query := fmt.Sprintf("SELECT id.id, id.username, id.domain, fullname.first_name, fullname.middle_name, fullname.last_name, email.email FROM identity.id LEFT JOIN identity.fullname ON fullname.id = id.id LEFT JOIN identity.email ON email.id = id.id WHERE id.id = '%s' ORDER BY id ASC;", user_id)

	rows, err := infrastructure.DbId.Query(query)
	if err != nil {
		return model.UserInfo{}, err
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
			return model.UserInfo{}, err
		}

		userList = append(userList, user_info)
	}

	if len(userList) != 1 {
		return model.UserInfo{}, fmt.Errorf("User not found")
	}

	return userList[0], nil
}
