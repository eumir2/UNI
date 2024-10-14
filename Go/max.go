package main
import "fmt"

/*Scrivere un programma num_max.go che legga una sequenza di 10 interi positivi
e stampi il massimo intero letto e quante volte il massimo compare nella sequenza*/

func main()  {
  var max int
  var c = 1
  var tmp int

  for i := 0; i < 10; i++{
    fmt.Scan(&tmp)
    if tmp > max{
      max = tmp
          c = 1
    }else if tmp == max{
      c++
    }
  }

  fmt.Printf("max: %d  volte: %d", max,c)
}
