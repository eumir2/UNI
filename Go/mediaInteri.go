package main
import "fmt"

//programma che calcola la media di un numero n di valori presi in input

func main()  {
  //varibili
  var n, i int
  var somma, tmp ,media float64

  fmt.Print("Inserire il numero di numeri: ")
  fmt.Scan(&n)

  for i = 0; i < n; i++{
    fmt.Print("(",i+1,"/",n,")Inserire un numero:")
    fmt.Scan(&tmp)
    //fmt.Println()

    somma += tmp
  }

  media = (somma / float64(n))
  fmt.Println("Media:", media)
}
