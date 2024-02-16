package userRouter

import (
	"testePostgres/user/controller"
	"github.com/gofiber/fiber/v2"
)

func NewRouter () *fiber.App {
  router := fiber.New()
  router.Get("/", userController.HandleGetAll)
  router.Post("/", userController.HandleCreateUser)
  return router
}
