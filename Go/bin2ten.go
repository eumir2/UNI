package main
import "fmt"
import "math"

/*Scrivere un programma bin2ten.go che converte un intero composto di soli 0 e 1 nel valore corrispondente al numero binario rappresentato (es. 101 -> 5).
 Nel caso il numero contenga altre cifre, il programma stampa un messaggio di errore.*/

func main()  {
  var bin int
  var dec int

  fmt.Scan(&bin)

  for i := 0; bin != 0; i++ {
    tmp :=  bin % 10
    if tmp == 1 {
      dec += tmp * int(math.Pow(2,float64(i)))
    }
    bin /= 10
  }
  fmt.Println(dec)
}
