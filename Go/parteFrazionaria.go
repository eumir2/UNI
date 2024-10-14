package main
import "fmt"

//programma che, dato un numero con la virgola, ne restituisce la parte frazionaria

func  main()  {
  //variabili
  var num float64
  var tmp int

  fmt.Print("Inserire un numero con la virgola: ")
  fmt.Scan(&num)

  tmp = int(num)

  fmt.Println("Parte frazionaria:", (num - float64(tmp)))
}
