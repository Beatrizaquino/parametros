package main

import (
	"fmt"
)

func maior(a int, b int) int {
	if a > b {
    return a
  } else {
    return b
  }
}

func repetir(txt string, n int) {
  for i:=0; i < n; i++ {
    fmt.Println(txt)
  }
  
  
}

func prmt() {
  var a = 5
  var b = 7
  var res = maior(a,b)

  fmt.Printf("O maior entre %d e %d é %d", a,b,res)
  repetir("Teste",10)
}
