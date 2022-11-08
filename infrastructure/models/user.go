package models

type User struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	BirthDate string `json:"birth_date"`
	Age       int    `json:"age"`
	Sexo      string `json:"sexo"`
}
