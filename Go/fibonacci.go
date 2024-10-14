package main
import "fmt"

/*Scrivere un programma fibonacci.go che legge un intero
positivo n e stampa un istogramma (orizzontale) dei numeri di
fibonacci dal primo all'n-esimo.*/


func main()  {
  var n int
  var val1 = 1
  var val2 = 1
  var somma int
  fmt.Scan(&n)

  fmt.Println("*")
  fmt.Println("*")

  for i:= 2 ; i < n; i++{
    somma = val1+val2
    for  j := 0; j < somma; j++{
      fmt.Print("*")
    }
    fmt.Println()
    val1,val2 = val2,somma
  }
}
