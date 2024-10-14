package main
import "fmt"

/*Scrivere un programma andamento.go che legge da tastiera una serie
(di almeno un elemento) di numeri interi > -1 e stampa "+" ogni volta che
il nuovo valore e` maggiore o uguale al precedente e "-" altrimenti.
Si ferma quando il numero in input e` -1 e stampa la somma di tutti i numeri letti.*/

func main()  {
  var tmp = 0
  var somma int
  var andamento string = ""

  fmt.Scan(&tmp)

  for tmp != -1 {
    c := tmp
    somma += tmp

    fmt.Scan(&tmp)

    if tmp != -1 {
      if c < tmp{
        andamento += "+"
      }else{
        andamento += "-"
      }
    }
  }

  fmt.Println(andamento)
  fmt.Println("somma:",somma)
}
