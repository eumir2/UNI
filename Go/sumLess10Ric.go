package main

import "fmt"

/*he, ricevuto un numero, restituisce la somma delle cifre del numero
se questa è minore di 10
o altrimenti il risultato della ri-applicazione della funzione
sulla somma delle cifre del numero.*/

func sumLess10Ric(n int) int {
  //caso base
  if n == 0{
    return 0
  }

  //ricorsione
  if n%10 + sumLess10Ric(n/10) <= 10{
    return  n%10 + sumLess10Ric(n/10)
  }else{
    return sumLess10Ric(n%10 + sumLess10Ric(n/10))
  }
}

func main()  {
  var  n int

  fmt.Scan(&n)
  fmt.Println(sumLess10Ric(n))


}
