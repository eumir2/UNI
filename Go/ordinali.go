package main

import "fmt"

/*Scrivere un programma ordinali.go dotato di una funzione ricorsiva

vonNeumann(n int) []byte

che, dato un numero naturale n, restituisca l'intero di von Neumann n,
e di una funzione main() per testarla.*/
func vonNeumann(n int) []byte{
  //caso base
  if n == 0{
    return []byte{}
  }else {//b = append(b,vonNeumann(n-1))
    return  append(vonNeumann(n-1),byte(n-1))
  }
}

func main()  {
  var n int

  fmt.Scan(&n)
  i := vonNeumann(n)
  fmt.Println(i)
}
