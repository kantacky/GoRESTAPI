package model

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Domain   string `json:"domain"`
}

type Password struct {
	Id           string `json:"id"`
	PasswordHash string `json:"password"`
}

type Fullname struct {
	Id         string `json:"id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}

type Email struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}
