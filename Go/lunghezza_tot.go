package main
import "fmt"

/*Scrivere un programma lunghezza_tot.go che legge da standard input
un int totLen e una sequenza di stringhe (una per riga) e
ne somma le lunghezze fino a raggiungere (o superare) totLen.
Quanto termina, stampa la somma delle lunghezze delle stringhe lette*/

func main()  {
  var totLen, c int
  var tmp string

  fmt.Scan(&totLen)

  for {
    fmt.Scan(&tmp)
    c+= len(tmp)

    if c >= totLen{
      break
    }
  }
  fmt.Println(c)
}
