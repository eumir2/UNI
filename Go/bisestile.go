package main
import "fmt"

/*Scrivere un programma bisestile.go che legge un intero n corrispondente all’anno di una
data, e determina se l’anno è bisestile o no (stampa "bisestile" o "non bisestile"). Vedi sotto la regola
per determinare se un anno è bisestile*/

func main()  {
  var anno int

  fmt.Print("anno: ")
  fmt.Scan(&anno)

  if anno%400 == 0 || anno%4 == 0 && anno%100 != 0{
    fmt.Println("bisestile")
  }else{
      fmt.Println("non bisestile")
  }
}
