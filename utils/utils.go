package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/klassmann/cpfcnpj"
)

func ValidateCPF (cpf int64) bool {
  CPFString := PadLeft(fmt.Sprint(cpf), 11)
  return cpfcnpj.ValidateCPF(CPFString)
}

func PadLeft(str string, length int) string {
  for len(str) < length {
    str = "0" + str
  }
  return str
}

func LoadConf () {
  err := godotenv.Load()
  if err != nil {
    fmt.Printf("unable to load environment\n")
    os.Exit(1)
  }
}

