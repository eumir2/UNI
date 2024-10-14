package main
import "fmt"

/*Specifiche: Scrivere un programma fizz_buzz.go che riceve in ingresso un numero intero e stampa
"Fizz" se il numero è multiplo di 3, "Buzz" se il numero è multiplo di 5, "Fizz Buzz" se è multiplo
sia di 3 sia di 5, niente altrimenti.*/

func main()  {
  var n, c int

  fmt.Print("numero: ")
  fmt.Scan(&n)

  if n % 3 == 0 && n % 5 == 0 {
    fmt.Println("Fizz Buzz")
    c++
  }

  if c == 0{
   if n % 3 == 0{
      fmt.Println("Fizz")
    }
    if n % 5 == 0 {
      fmt.Println("Buzz")
    }
  }
}
