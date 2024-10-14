package main
import "fmt"

/*Scrivere un programma Go consumo_resa.go che calcola il consumo medio e la
resa di un motore date la distanza totale percorsa (in km) e la quantità di carburante utilizzata
(in litri). I valori in input sono di tipo float64.*/

func main()  {
  var km, litri float64

  fmt.Print("distanza percorsa (in km): ")
  fmt.Scan(&km)

  fmt.Print("quantità di carburante utilizzata (in l): ")
  fmt.Scan(&litri)

  fmt.Println("consumo medio:", (litri / km), "l/km")
  fmt.Println("resa media:", (km / litri), "km/l")
}
