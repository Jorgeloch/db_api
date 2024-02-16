package userController

import (
  "fmt"
  "testePostgres/utils"
  "testePostgres/user/model"
  "github.com/gofiber/fiber/v2"
)

func HandleGetAll (c *fiber.Ctx) error{
  users, err := userModel.GetAll() 
  if err != nil {
    fmt.Printf("Unable to get user: %v\n", err)
    c.SendStatus(400)
  }
  return c.JSON(users)
}

func HandleCreateUser (c *fiber.Ctx) error{
  newUser := new(userModel.User)
  err := c.BodyParser(newUser) 
  if err != nil {
    fmt.Printf("unable to get user from request body: %v\n", err)
    return c.SendStatus(400)
  }
  err = userModel.ValidateUser(newUser)
  if err != nil {
    fmt.Printf("invalid user given by the request body\n")
    return c.SendStatus(400)
  }
  if !utils.ValidateCPF(newUser.CPF) {
    fmt.Printf("invalid cpf number\n")
  }
  err = userModel.Create(newUser)
  if err != nil {
    fmt.Printf("Unable to insert new user: %v\n", err)
    return c.SendStatus(500)
  }
  c.Status(201)
  return c.JSON(newUser)
}
