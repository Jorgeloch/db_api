package main

import (
	"github.com/jorgeloch/db_api/user/router"
	"github.com/jorgeloch/db_api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
  utils.LoadConf()
  app := fiber.New()
  app.Use(logger.New())
  app.Mount("/user", userRouter.NewRouter())
  app.Listen(":5050")
}

