package main
import "fmt"

/*Scrivere un programma crescente.go che legge da standard input una stringa
di sole lettere minuscole e la stampa inserendo un'-' ogni volta
che una lettera viene prima in ordine alfabetico della lettera
precedente.*/

func main()  {
  var s string

  fmt.Scan(&s)
  tmp := s[0]

  for i := 0; i < len(s); i++{
    if(s[i] < tmp){
      fmt.Print("-")
    }
    tmp = s[i]
    fmt.Print(string(s[i]))
  }
}
