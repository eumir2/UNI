package main

import "fmt"
import "os"
import "strconv"

/*Scrivere un programma potenza.go, dotato di una funzione
ricorsiva

func potenza(base, exp int) int

che restituisce il valore della base elevata a exp
*/
func potenza(b, e int)int{
  if e == 1{
    return b
  }else{
    return b * potenza(b,e-1)
  }
}

func main()  {
  base,_ := strconv.Atoi(os.Args[1])
  esponente,_ := strconv.Atoi(os.Args[2])

  fmt.Println(potenza(base,esponente))
}
