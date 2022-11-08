package users

import (
	"github.com/jmoiron/sqlx"

	"estrellitas-crud/infrastructure/logger"
	"estrellitas-crud/infrastructure/models"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesUsersRepository interface {
	create(m *Users) error
	update(m *Users) error
	delete(id string) error
	getByID(id string) (*Users, error)
	getAll() ([]*Users, error)
	getByEmail(email string) (*Users, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesUsersRepository {
	var s ServicesUsersRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newUsersPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
