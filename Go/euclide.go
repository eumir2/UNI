package main
import "fmt"

/*Scrivere un programma euclide.go che legge da standard input due interi
a e b, con a >= b, e calcola il MCD tra i due numeri con l'algoritmo di Euclide.*/

func main()  {
  var a,b int

  fmt.Print("a,b: ")
  fmt.Scan(&a, &b)

  if b == 0{
    fmt.Println(a)
  }else{
    for b != 0{
      r := a % b
      a = b
      b = r
    }
    fmt.Println(a)
  }
}
