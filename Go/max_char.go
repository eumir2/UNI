package main
import "fmt"

/*Scrivere un programma max_char.go che legge da standard input
una sequenza di 5 caratteri ASCII (byte) e stampa il maggiore
in ordine lessicografico (codice ASCII).*/

func main()  {
  var carattere,a_capo byte
  var max int
  max = 0


  for i := 0; i < 5; i++{
    fmt.Scanf("%c %c", &carattere, &a_capo)
    //fmt.Printf("%c \n",carattere)
    if int(carattere) > max{
        max = int(carattere)
    }
  }
  fmt.Printf("%c \n", max)
}
