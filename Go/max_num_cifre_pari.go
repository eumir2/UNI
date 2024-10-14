package main
import "fmt"
import "strconv"

/*Scrivere un programma max_num_cifre_pari.go che data una sequenza di numeri
(letti come stringhe), stampa il massimo numero di cifre pari contenute in un numero.*/

func main()  {
  var s string
  var c int

  fmt.Scan(&s)
  for  _,r := range s {
    i,_ := strconv.Atoi(string(r))
    if i % 2 == 0{
      c++
    }
  }
  fmt.Println(c)
}
