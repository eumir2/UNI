package main
import "fmt"

/*Scrivere un programma capitale.go che legge da stardard input
tre valori float64 che rappresentano un capitale iniziale (es. 1000.50 euro),
un tasso di interesse annuale (es. 1.3%) e un obiettivo (es. 2000 euro),
e calcola quanti anni occorrono per arrivare a (o superare) l'obiettivo.*/

func main()  {
  var c_i, t_i, o float64
  var anni int
  fmt.Print("capitale iniziale: ")
  fmt.Scan(&c_i)

  fmt.Print("tasso di interesse: ")
  fmt.Scan(&t_i)

  fmt.Print("obiettivo: ")
  fmt.Scan(&o)

  c_tmp := c_i //capitale temporaneo = capitale iniziale

  for ;  c_tmp <= o; {
    c_tmp += t_i * c_tmp
    //fmt.Println(c_tmp)
    anni++
  }

  fmt.Println(anni)

}
