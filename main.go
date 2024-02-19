package main

import (
	_ "github.com/jorgeloch/db_api/docs"
	"github.com/jorgeloch/db_api/user/router"
	"github.com/jorgeloch/db_api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

//	@title			Databases College Homeowrk API
//	@version		1.0
//	@description	This is a API created as a homework to COMP0455 class

func main() {
	utils.LoadConf()
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Mount("/user", userRouter.NewRouter())
	app.Listen(":5050")
}
