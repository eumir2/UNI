package main
import "fmt"

/*Scrivere un programma conta_vocali.go che legge da standard input
una parola e stampa il numero di lettere (rune) che sono vocali (aeiou).*/

func main()  {
  var s string
  var v int

  fmt.Scan(&s)

  for i := 0; i < len(s); i++{
    if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u'{
      v++
    }
  }
  fmt.Println(v)
}
