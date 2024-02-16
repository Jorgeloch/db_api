package userModel

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
)

type User struct {
  Name            string      `validate:"required,max=40" json:"name" db:"name"`
  CPF             int64       `validate:"required" json:"cpf" db:"cpf"`
  DataNascimento  time.Time   `validate:"required" json:"data_nascimento" db:"data_nascimento"`
}

type database interface {
  GetAll() ([]User, error)
  Create(newUser *User) error
}

type DatabaseConnection struct {
  conn *pgx.Conn
} 

func ConnectDatabase() *DatabaseConnection {
  URL := os.Getenv("DATABASE_URL")
  conn, err := pgx.Connect(context.Background(), URL)
  if err != nil {
    fmt.Printf("Unable to connect to database: %v\n", err)
    os.Exit(1)
  } 
  return &DatabaseConnection{conn: conn}
}

func GetAll() ([]User, error) {
  db := ConnectDatabase()
  rows, _ := db.conn.Query(context.Background(), 
    `
    SELECT name, cpf, data_nascimento 
    FROM users
    `)

  users, err := pgx.CollectRows(rows, pgx.RowToStructByName[User]) 
  if err != nil {
    fmt.Printf("Unable to convert user: %v\n", err)
    db.conn.Close(context.Background())
    return nil, err
  }
  db.conn.Close(context.Background())
  return users, nil
}

func Create(newUser *User) error {
  db := ConnectDatabase()
  args := pgx.NamedArgs{
    "name": newUser.Name,
    "cpf": newUser.CPF,
    "dataNascimento": newUser.DataNascimento,
  }
  _, err := db.conn.Exec(context.Background(), `
    INSERT INTO users
    (name, cpf, data_nascimento) VALUES (@name, @cpf, @dataNascimento)
    `, args) 
  db.conn.Close(context.Background())
  return err
}

func ValidateUser (user *User) error {
  validate := validator.New()
  return validate.Struct(user)
}
