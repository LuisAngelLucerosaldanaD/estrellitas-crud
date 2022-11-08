package users

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Model estructura de Users
type Users struct {
	ID        string    `json:"id" db:"id" valid:"required,uuid"`
	Name      string    `json:"name" db:"name" valid:"required"`
	Lastname  string    `json:"lastname" db:"lastname" valid:"required"`
	Email     string    `json:"email" db:"email" valid:"required"`
	BirthDate string    `json:"birth_date" db:"birth_date" valid:"required"`
	Sexo      string    `json:"sexo" db:"sexo" valid:"required"`
	Age       int       `json:"age" db:"age" valid:"required"`
	Password  string    `json:"password" db:"password" valid:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewUsers(id string, name string, lastname string, email string, birthDate string, sexo string, age int, password string) *Users {
	return &Users{
		ID:        id,
		Name:      name,
		Lastname:  lastname,
		Email:     email,
		BirthDate: birthDate,
		Sexo:      sexo,
		Age:       age,
		Password:  password,
	}
}

func (m *Users) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
