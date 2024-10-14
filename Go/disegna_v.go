package main
import "fmt"

/*Scrivere un programma disegna_v.go che legge un intero
positivo n e stampa una v di asterischi di altezza n.*/

func main()  {
  var n int
  fmt.Scan(&n)

 //si = spazi iniziali
 //sm = spazi nel mezzo

  for i := 0; i < n; i++{
    si := 0
    for si = 0; si <= i; si++{
      fmt.Print(" ")
    }

    fmt.Print("*")

    sm := n - i
    if sm > i {
      for sm := n - i ; sm > i; sm--{
         fmt.Print(" ")
       }
      fmt.Println("*")
    }else{
      break
    }
  }
}
