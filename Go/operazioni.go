package main

import "fmt"

/*func operazioni(n1,n2 int) (int,int,int) {
  somma := n1 + n2
  prodotto := n1 * n2
  differenza := n1 - n2
  return somma,prodotto,differenza
}*/

func operazioni(n1,n2 int) (somma,prodotto,differenza int) {
  somma = n1 + n2
  prodotto = n1 * n2
  differenza = n1 - n2
  return
}


func main()  {
var n1,n2 int

  fmt.Scan(&n1,&n2)

  s,p,d := operazioni(n1,n2)

  fmt.Printf("%4d %4d %4d ",s,p,d)
}
