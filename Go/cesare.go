package main
import "fmt"

/*Scrivere un programma cesare.go che legge da standard input
un valore intero non negativo k (la chiave di cifratura) e
una sequenza di lettere minuscole consecutive
(sulla stessa riga e senza spazi) terminate da <invio> ('\n').
Il programma stampa la sequenza letta cifrata secondo il
cifrario di Cesare, usando come chiave k (quella fornita dall'utente):
ogni lettera del testo in chiaro è sostituita nel
testo cifrato dalla lettera che si trova k posizioni dopo
nell'alfabeto, ritornando dopo la zeta alla lettera a.*/

func main()  {
  var s string
  var k int
  var c string

  fmt.Print("chiave: ")
  fmt.Scan(&k)

  fmt.Scan(&c)

  for i := 0; i < len(c); i++{
    s += string(c[i] + byte(k))
  }

  fmt.Println(s)
}
