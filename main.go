package main

import (
	"estrellitas-crud/api"
	"estrellitas-crud/infrastructure/env"
)

// @title Estrellitas crud
// @version 1.0
// @description Documentation for estrellitas crud
// @termsOfService https://www.estrellitas.net/terms/
// @contact.name API Support
// @contact.email info@estrellitas.net
// @license.name Software Owner
// @license.url https://www.estrellitas.net/terms/licenses
// @host http://127.0.0.1:50050
// @tag.name User
// @tag.description Methods of user
// @BasePath /
func main() {
	c := env.NewConfiguration()

	api.Start(c.App.Port, c.App.ServiceName, c.App.LoggerHttp, c.App.AllowedDomains)
}
