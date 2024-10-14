package main
import "fmt"

/*: Scrivere un programma tasse.go che chiede reddito (int) e se coniugato (bool)
e stampa le tasse da pagare secondo la seguente tabella.
Usate costanti per le aliquote (10% e 25%) e per gli scaglioni (32000 e 64000).

non coniugato     coniugato     %
0 - 32000         0 - 64000    10%
 > 32000          > 64000      25*/

func main()  {
  const a1 = 10
  const a2 = 25
  const con1 = 32000
  const con2 = 64000

  var reddito, tasse int
  var coniugato bool

  fmt.Print("reddito: ")
  fmt.Scan(&reddito)

  fmt.Print("coniugato (t/f): ")
  fmt.Scan(&coniugato)

  if coniugato == true{
    if reddito >= 0 && reddito <= con1 {
      tasse = reddito / 100 * a1
    }else{
      tasse = reddito / 100 * a2
    }
  }else{
    if reddito >= 0 && reddito <= con2{
      tasse = reddito / 100 * a1
    }else{
      tasse = reddito / 100 * a2
    }
  }

  fmt.Println("tasse:",tasse)


}
