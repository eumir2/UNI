package main
import "fmt"

/*Scrivere un programma indovina_10.go che fissa nel programma un numero intero (a
vostra scelta) tra 1 e 10 da indovinare, legge un intero da standard input e
• se il numero in input è fuori dall’intervallo 1-10, stampa “Valore non valido”;
• se il numero è quello fissato, stampa “Hai indovinato!”;
• altrimenti stampa “Non hai indovinato”.*/

func main()  {
  const giusto = 2
  var n int

  fmt.Scan(&n)

  if n < 1 || n > 10 {
    fmt.Println("Valore non valido")
  }else if n == giusto  {
    fmt.Println("Hai indovinato")
  }else{
    fmt.Println("Non hai indovinato")
  }
}
