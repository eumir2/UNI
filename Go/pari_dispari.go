package main
import "fmt"

/*Scrivere un programma pari_dispari.go che legge un intero n e, a seconda del valore di
n, stampa uno dei messaggi "n è pari" oppure "n è dispari"*/

func main()  {
  var n int

  fmt.Print("numero: ")
  fmt.Scan(&n)

  if n % 2 == 0{
    fmt.Println(n, "è pari")
  }else{
    fmt.Println(n, "è dispari")
  }
}
