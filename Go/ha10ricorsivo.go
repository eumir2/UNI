package main

import "fmt"

/*Scrivere una funzione ricorsiva ha10Ricorsiva che,
avendo in input una lista di interi,
dia in output TRUE se 10 è un elemento della lista,
 FALSE altrimenti.*/

func ha10(s []int)bool  {
  l := len(s)

  if l == 0{
    return false
  }

  if s[0] == 10{
    return true
  }else{
    return ha10(s[1:])
  }
}

 func main()  {
   var s = []int{2,3,5,3,6,4,43,12,54}
   //fmt.Println(s)

   fmt.Println(ha10(s))
 }
