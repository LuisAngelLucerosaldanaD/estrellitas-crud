package users

import (
	"fmt"
	"github.com/asaskevich/govalidator"

	"estrellitas-crud/infrastructure/logger"
	"estrellitas-crud/infrastructure/models"
)

type PortsServerUsers interface {
	CreateUsers(id string, name string, lastname string, email string, birthDate string, sexo string, age int, password string) (*Users, int, error)
	UpdateUsers(id string, name string, lastname string, email string, birthDate string, sexo string, age int, password string) (*Users, int, error)
	DeleteUsers(id string) (int, error)
	GetUsersByID(id string) (*Users, int, error)
	GetAllUsers() ([]*Users, error)
	GetUsersByEmail(email string) (*Users, int, error)
}

type service struct {
	repository ServicesUsersRepository
	user       *models.User
	txID       string
}

func NewUsersService(repository ServicesUsersRepository, user *models.User, TxID string) PortsServerUsers {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateUsers(id string, name string, lastname string, email string, birthDate string, sexo string, age int, password string) (*Users, int, error) {
	m := NewUsers(id, name, lastname, email, birthDate, sexo, age, password)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Users :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateUsers(id string, name string, lastname string, email string, birthDate string, sexo string, age int, password string) (*Users, int, error) {
	m := NewUsers(id, name, lastname, email, birthDate, sexo, age, password)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Users :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteUsers(id string) (int, error) {
	if !govalidator.IsUUID(id) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id isn't uuid"))
		return 15, fmt.Errorf("id isn't uuid")
	}

	if err := s.repository.delete(id); err != nil {
		if err.Error() == "ecatch:108" {
			return 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't update row:", err)
		return 20, err
	}
	return 28, nil
}

func (s *service) GetUsersByID(id string) (*Users, int, error) {
	if !govalidator.IsUUID(id) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id isn't uuid"))
		return nil, 15, fmt.Errorf("id isn't uuid")
	}
	m, err := s.repository.getByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAllUsers() ([]*Users, error) {
	return s.repository.getAll()
}

func (s *service) GetUsersByEmail(email string) (*Users, int, error) {
	if !govalidator.IsEmail(email) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("email isn't email valid"))
		return nil, 15, fmt.Errorf("email isn't email valid")
	}
	m, err := s.repository.getByEmail(email)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByEmail row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}
