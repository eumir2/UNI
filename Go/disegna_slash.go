package main
import "fmt"

/*Scrivere un programma disegna_slash.go che legge un intero
positivo n e stampa uno slash (\) di asterischi di altezza n.*/

func main()  {
  var n int

  fmt.Scan(&n)

  for i := 0; i < n; i++{
    for j := 0; j < i; j++{
      fmt.Print(" ")
    }
    fmt.Println("\\")
  }
}
