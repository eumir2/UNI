package main
import "fmt"

/*Scrivere un programma pos_neg_nullo.go che legge un int da standard input, stabilisce
se è positivo, negativo o nullo e stampa il messaggio corrispondente*/

func main()  {
  var n int

  fmt.Print("un int: ")
  fmt.Scan(&n)

  if n > 0 {
    fmt.Println("positivo")
  }else if n < 0 {
    fmt.Println("negativo")
  }else{
    fmt.Println("nullo")
  }
}
