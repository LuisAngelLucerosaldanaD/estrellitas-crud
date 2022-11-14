package users

import (
	"estrellitas-crud/infrastructure/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterUser(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerUser{DB: db, TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/user")
	user.Post("/create", h.createUser)
	user.Post("/login", h.Login)
	user.Post("/delete", middleware.JWTProtected(), h.Delete)
	user.Get("/:id", middleware.JWTProtected(), h.getUserById)
}
