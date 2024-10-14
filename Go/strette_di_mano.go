package main

import "fmt"
import "os"
import "strconv"


/*Scrivere un programma strette_di_mano.go, dotato di una funzione
ricorsiva

numStretteMano(n int) int

che calcola il numero di strette di mano scambiate nella stanza se ciascuna
persona stringe la mano a ciascuna altra una e una sola volta,
*/

func numStretteMano(n int) int{
  if n == 0{
    return 0
  }else{
    //fmt.Println(n)
    return n + numStretteMano(n - 1)
  }
}

func main()  {
  strette,_ := strconv.Atoi(os.Args[1])

  fmt.Println(numStretteMano(strette-1))
}
