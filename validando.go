package main

import (
	"fmt"
	"strings"
)

func main() {
  var email string
  fmt.Println("Insira um EMail: ")
  fmt.Scan(&email)
  var valido bool = validarEmail(email)
  if valido {
    fmt.Println("Emaiol validado")
  } else {
    fmt.Println("Email invalido")
  }

}

func validarEmail(email string) bool {
  if !strings.Contains(email, "@") { return false } 
  if !strings.Contains(email, ".") { return false }
  if len(email) < 3 { return false}

    return true

}



