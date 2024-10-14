package main
import "fmt"
import "math/rand"
import "time"

/*Scrivere un programma somma.go che genera 10 numeri interi casuali
tra 0 e 10, li stampa, e stampa la somma dei valori generati.*/

func main()  {
  somma := 0
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  for i := 0; i < 10; i++{
    somma += r1.Intn(10)
    fmt.Println(somma)
  }
  //fmt.Println(somma)

}
