package main
import "fmt"

/* Scrivere un programma primo.go che, dato un numero intero su standard input,
determina se il numero e` primo.
Suggerimento: occorre determinare il primo numero che e` un divisore (se c'è).
Domanda: dato in input n, quante iterazioni dovrò fare al massimo?*/


func main()  {
  var primo int
  var d int
  fmt.Scan(&primo)

  for d = 2; d <= primo/2; d++{
    if primo % d == 0{
      break
    }
  }

  if d < primo / 2 {
    fmt.Println("non primo")
  }else{
    fmt.Println("primo")
  }
}
