package main
import "fmt"

/*: Scrivere un programma punto_retta.go che legge da standard input le cooordinate
(x1,y1) di un punto P e i coefficienti m e q di una retta y = mx + q; controlla se il punto P sta sopra,
sotto o appartiene alla retta y data, e stampa rispettivamente "sopra", "sotto", "sulla retta".
Si dichiarino i dati in ingresso come float64.
Nota. Trattando con numeri float, cosideriamo appartenenti alla retta i punti che distano dalla retta
meno di 10^-6. In Go 10^-6 può essere scritto come 1e-06*/

func main()  {
  var x1, y1, m, q float64

  fmt.Print("x e y del punto: ")
  fmt.Scan(&x1, &y1)

  fmt.Print("m e q della retta: ")
  fmt.Scan(&m, &q)

  tmp := y1 - (x1 * m) - q

  if tmp < 0{
    fmt.Println("sotto")
  }else if tmp > 0 {
    fmt.Println("sopra")
  }else{
    fmt.Println("sulla retta")
  }

}
