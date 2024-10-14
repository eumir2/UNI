package main
import "fmt"

/*Scrivere un programma Go intersezione_rette.go che, date le equazioni di due
rette, stabilisce in che punto si intersecano. I dati in input sono di tipo float64.
Nota. L’equazione della retta è: y = mx + q. Dal punto di vista matematico, occorre risolvere
un sistema di primo grado.
Esempio di esecuzione:
retta 1: m e q? 1 4
retta 2: m e q? 2 6
intersezione in (-2,2)*/

func main()  {
  var m1, q1, m2, q2 float64

  fmt.Print("retta 1: m e q? ")
  fmt.Scan(&m1, &q1)

  fmt.Print("retta 2: m e q? ")
  fmt.Scan(&m2, &q2)

  //x
  x := (q2 - q1) / (m1 - m2)

  //y
  y := m1 * x + q1

  fmt.Println("intersezione in (", x, ",", y, ")")
}
