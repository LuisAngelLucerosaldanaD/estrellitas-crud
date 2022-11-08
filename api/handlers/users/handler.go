package users

import (
	"estrellitas-crud/infrastructure/jwt"
	"estrellitas-crud/infrastructure/logger"
	"estrellitas-crud/infrastructure/models"
	"estrellitas-crud/infrastructure/msg"
	"estrellitas-crud/infrastructure/pwd"
	"estrellitas-crud/pkg/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerUser struct {
	DB   *sqlx.DB
	TxID string
}

// createUser godoc
// @Summary Create User
// @Description Create User
// @tags User
// @Accept  json
// @Produce  json
// @Param createUser body requestCreateUser true "Request create user"
// @Success 200 {object} responseCreateUser
// @Router /api/v1/user/create [post]
func (h *handlerUser) createUser(c *fiber.Ctx) error {
	res := responseCreateUser{Error: true}
	m := requestCreateUser{}
	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf("couldn't bind model create wallets: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if m.Password != m.ConfirmPassword {
		logger.Error.Printf("this password is not equal to confirm_password")
		res.Code, res.Type, res.Msg = msg.GetByCode(10005, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvAuth := auth.NewServerAuth(h.DB, nil, h.TxID)

	_, code, err := srvAuth.SrvUser.CreateUsers(uuid.New().String(), m.Name, m.Lastname, m.Email, m.BirthDate, m.Sexo, m.Age, m.Password)
	if err != nil {
		logger.Error.Printf("couldn't create user, error: " + err.Error())
		res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = "Usuario creado correctamente"
	res.Code, res.Type, res.Msg = msg.GetByCode(29, h.DB, h.TxID)
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

// getUserById godoc
// @Summary get user by ID
// @Description get user by ID
// @tags User
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization" default(Bearer <Add access token here>)
// @Param id path string true "user ID"
// @Success 200 {object} responseUser
// @Router /api/v1/user/{id} [get]
func (h *handlerUser) getUserById(c *fiber.Ctx) error {
	res := responseUser{Error: true}
	usrId := c.Params("id")
	srvAuth := auth.NewServerAuth(h.DB, nil, h.TxID)

	user, code, err := srvAuth.SrvUser.GetUsersByID(usrId)
	if err != nil {
		logger.Error.Printf("couldn't get user by id, error: " + err.Error())
		res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if user == nil {
		res.Code, res.Type, res.Msg = msg.GetByCode(5, h.DB, h.TxID)
		res.Error = false
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = &models.User{
		Email:     user.Email,
		Name:      user.Name,
		Lastname:  user.Lastname,
		BirthDate: user.BirthDate,
		Age:       user.Age,
		Sexo:      user.Sexo,
	}
	res.Code, res.Type, res.Msg = msg.GetByCode(29, h.DB, h.TxID)
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

// Login godoc
// @Summary Login Estrellitas Crud
// @Description Login Estrellitas Crud
// @tags User
// @Accept  json
// @Produce  json
// @Param Login body RequestLogin true "Request login"
// @Success 200 {object} ResponseLogin
// @Router /api/v1/user/login [post]
func (h *handlerUser) Login(c *fiber.Ctx) error {

	res := ResponseLogin{Error: true}
	m := RequestLogin{}
	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf("couldn't bind model login: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvAuth := auth.NewServerAuth(h.DB, nil, h.TxID)

	usr, code, err := srvAuth.SrvUser.GetUsersByEmail(m.Email)
	if err != nil {
		logger.Error.Printf("couldn't get user by email: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if !pwd.Compare(usr.ID, usr.Password, m.Password) {
		logger.Error.Printf("Usuario o contraseña incorrectos")
		res.Code, res.Type, res.Msg = 22, 1, "Usuario o contraseña incorrectos"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	token, code, err := jwt.GenerateJWT(models.User{
		Id:        usr.ID,
		Email:     usr.Email,
		Name:      usr.Name,
		Lastname:  usr.Lastname,
		BirthDate: usr.BirthDate,
		Age:       usr.Age,
		Sexo:      usr.Sexo,
	})
	if err != nil {
		logger.Error.Printf("couldn't generate token: " + err.Error())
		res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	authRes := Token{AccessToken: token, RefreshToken: token}
	res.Data = authRes
	res.Code, res.Type, res.Msg = msg.GetByCode(29, h.DB, h.TxID)
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

// Delete godoc
// @Summary Delete Estrellitas Crud
// @Description Delete Estrellitas Crud
// @tags User
// @Accept  json
// @Produce  json
// @Param Delete body DeleteUser true "Request delete user"
// @Success 200 {object} ResponseDeleteUser
// @Router /api/v1/user/delete [post]
func (h *handlerUser) Delete(c *fiber.Ctx) error {

	res := ResponseDeleteUser{Error: true}
	m := DeleteUser{}
	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf("couldn't bind model login: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvAuth := auth.NewServerAuth(h.DB, nil, h.TxID)

	usr, code, err := srvAuth.SrvUser.GetUsersByEmail(m.Email)
	if err != nil {
		logger.Error.Printf("couldn't get user by email: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(code, h.DB, h.TxID)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if !pwd.Compare(usr.ID, usr.Password, m.Password) {
		logger.Error.Printf("Usuario o contraseña incorrectos")
		res.Code, res.Type, res.Msg = 22, 1, "Usuario o contraseña incorrectos"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	code, err = srvAuth.SrvUser.DeleteUsers(usr.ID)
	if err != nil {
		logger.Error.Printf("Couldn't delete user")
		res.Code, res.Type, res.Msg = 22, 1, "No se pudo eliminar al usuario"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Code, res.Type, res.Msg = msg.GetByCode(29, h.DB, h.TxID)
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}
