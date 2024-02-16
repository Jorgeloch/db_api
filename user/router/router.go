package userRouter

import (
	"github.com/jorgeloch/db_api/user/controller"
	"github.com/gofiber/fiber/v2"
)

func NewRouter () *fiber.App {
  router := fiber.New()
  router.Get("/", userController.HandleGetAll)
  router.Post("/", userController.HandleCreateUser)
  return router
}
