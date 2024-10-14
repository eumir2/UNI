package main
import "fmt"

/*Scrivere un programma differenze.go che legge una serie di valori
float64 da tastiera e stampa le differenze, cioe` la differenza
tra il secondo e il primo, tra il terzo e il secondo, e cosi' via.
Il programma termina quando riceve in input il valore 0.*/

func main()  {
  var n = 1
  tmp := 0
  for n != 0 {
    fmt.Scan(&n)
    differenza := n - tmp
    fmt.Println(n,"-",tmp,":", differenza)
    tmp = n
  }
}
