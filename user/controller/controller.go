package userController

import (
  "fmt"
  "github.com/jorgeloch/db_api/utils"
  "github.com/jorgeloch/db_api/user/model"
  "github.com/gofiber/fiber/v2"
)

//  GetAllUsers godoc
//	@Summary		Get all users
//	@Description Get all users	
//	@Produce		json
//	@Success		200	{array}	userModel.User
//	@Failure		400	
//	@Failure		500
//	@Router			/user [get]
func HandleGetAll (c *fiber.Ctx) error{
  users, err := userModel.GetAll() 
  if err != nil {
    fmt.Printf("Unable to get user: %v\n", err)
    c.SendStatus(400)
  }
  return c.JSON(users)
}

// CreateUser godoc
//	@Summary		Create a new User
//	@Description	Create a new User
//	@Accept			json
//	@Produce		json
//  @Param user body userModel.User true "user model to be created" Format(json)
//	@Success		200	{object}	userModel.User
//	@Failure		400	
//	@Failure		500
//	@Router			/user [post]
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
