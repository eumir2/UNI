package main

import "fmt"

/*programma che, dati due numeri naturali, restituisca il valore del loro
massimo comun divisore
Nota: Il MCD di due interi positivi e` il piu' grande intero che li
divide entrambi.*/
func mcd(p, q int)int  {
  if p > q && q == 0{
    return p
  }else{
    return mcd(q,p%q)
  }
}



func main()  {
  var p,q int

  fmt.Scan(&p, &q)
  fmt.Println(mcd(p,q))
}
