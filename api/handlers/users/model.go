package users

import "estrellitas-crud/infrastructure/models"

type requestCreateUser struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Name            string `json:"name"`
	Lastname        string `json:"lastname"`
	BirthDate       string `json:"birth_date"`
	Age             int    `json:"age"`
	Sexo            string `json:"sexo"`
}

type responseCreateUser struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
	Code  int    `json:"code"`
	Type  int    `json:"type"`
	Msg   string `json:"msg"`
}
type responseUser struct {
	Error bool         `json:"error"`
	Data  *models.User `json:"data"`
	Code  int          `json:"code"`
	Type  int          `json:"type"`
	Msg   string       `json:"msg"`
}

type ResponseLogin struct {
	Error bool   `json:"error"`
	Data  Token  `json:"data"`
	Code  int    `json:"code"`
	Type  int    `json:"type"`
	Msg   string `json:"msg"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DeleteUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseDeleteUser struct {
	Error bool   `json:"error"`
	Code  int    `json:"code"`
	Type  int    `json:"type"`
	Msg   string `json:"msg"`
}
