package main
import "fmt"

/*Scrivere un programma cifra_pari.go che, dato un intero da standard input, determina e stampa in che posizione (procedendo da destra a sinistra) si trova la prima cifra pari del numero.
Se il numero non contiene cifre pari, il programma stampa -1.*/

func main()  {
  var n int
  var c = 1 //contatore di cifre
  var f = 1 //flag di controllo

  fmt.Scan(&n)


  for n > 0 {
    tmp := n % 10
    if tmp % 2 == 0 {
      f = 0
      break
    }
    n /= 10
    c++
  }

  if f == 0{
    fmt.Println("prima cifra pari in posizione:", c)
  }else{
    fmt.Println(-1)
  }

}
