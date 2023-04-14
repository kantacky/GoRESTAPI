package auth

import (
	"kantacky.com/api/handler/v1/util"
	"kantacky.com/api/model"
)

func AuthUser(username string, domain string, password string) (model.TokenIssued, error) {
	passwordHash := util.GetHashValueOf(password)

	id, err := FetchUserId(username, domain)
	if err != nil {
		return model.TokenIssued{}, err
	}

	id, err = VerifyPassword(id, passwordHash)
	if err != nil {
		return model.TokenIssued{}, err
	}

	token, err := IssueToken(id)
	if err != nil {
		return model.TokenIssued{}, err
	}

	res := model.TokenIssued{
		Token: token,
	}

	return res, nil
}
