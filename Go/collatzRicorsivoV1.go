package main

import "fmt"

/*Scrivere una funzione ricorsiva collatz(n) che calcola la successione di Collatz
 generata a partire da un valore n>0,
 cioè la sequenza di numeri generati secondo la funzione f così definita:
f(1) =   1
f(n) =   n/2         se n è pari
	     n*3+1     se n è dispari*/
func collatz(n int)int{
  if n == 0{
    return 0
  }
  if n == 1{
    return 1
  }
  if n > 0 {
    if n % 2 == 0{
      return collatz(n/2)
    }else{
      return collatz(n*3+1)
    }
  }
  return n
}
func main()  {
  var n int
  fmt.Scan(&n)

  fmt.Println(collatz(n))
}
