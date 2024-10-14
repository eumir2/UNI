package main
import "fmt"

/*Scrivere un programma sovrapposizione.go che legge da tastiera il giorno [1-31], l'ora di
inizio [0-24] e l'ora di fine [0-24] di due appuntamenti e stabilisce se si sovrappongono (anche
parzialmente) o no.*/

func main()  {
  var g1, s1, e1 int
  var g2, s2, e2 int
  var c = 0 //controllo, 0 se non si sovrappongono
  fmt.Print("appuntamento 1 (gg, start, end):")
  fmt.Scan(&g1, &s1, &e1)

  fmt.Print("appuntamento 2 (gg, start, end):")
  fmt.Scan(&g2, &s2, &e2)

  if g1 == g2{
    if e1 > s2{
      c++
    }
  }
  if c == 0{
    fmt.Println("non si sovrappongono")
  }else{
    fmt.Println("si sovrappongono")
  }
}
