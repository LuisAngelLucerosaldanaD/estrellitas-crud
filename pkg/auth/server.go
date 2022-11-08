package auth

import (
	"estrellitas-crud/infrastructure/models"
	"estrellitas-crud/pkg/auth/users"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	SrvUser users.PortsServerUsers
}

func NewServerAuth(db *sqlx.DB, user *models.User, txID string) *Server {

	repoUser := users.FactoryStorage(db, user, txID)
	srvUser := users.NewUsersService(repoUser, user, txID)

	return &Server{
		SrvUser: srvUser,
	}
}
