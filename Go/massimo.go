package main
import "fmt"
import "math/rand"
import "time"

/*Scrivere un programma massimo.go che genera 10 numeri interi casuali
tra 10 e 30, li stampa, e stampa il massimo valore generato.
*/

func main()  {
  max := 0
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  for i := 0; i < 10; i++{
    tmp := r1.Intn(30 - 10 + 1) + 10
    fmt.Println(tmp)
    if tmp > max {
      max = tmp
    }
  }
  fmt.Println("max: ", max)

}
