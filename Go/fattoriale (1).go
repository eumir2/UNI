package main

import "fmt"
import "os"
import "strconv"

/*Scrivere un programma fattoriale.go, dotato di una
funzione ricorsiva */

func fattoriale(n int)int{
  if n == 0{
    return 1
  }else{
    return n * fattoriale(n-1)
  }
}

func main()  {
  var n int

  n,_ = strconv.Atoi(os.Args[1])

  fatt := fattoriale(n)
  fmt.Println(fatt)
}
