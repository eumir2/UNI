package main

import "fmt"

/*Scrivere una funzione ricorsiva quanteVolteRicorsiva che,
 data una lista di rune s ed una runa c,
 restituisca il numero di occorrenze di c in s.*/

func quanteVolteRicorsiva(s []rune, c rune)int  {
  cont := 0
  if len(s) == 0{
    return 0
  }
  if s[0] == c{
    cont++
  }
  return cont + quanteVolteRicorsiva(s[1:],c)

}


func main()  {
  var s = []rune{'a','c','f','c','t','p'}

  fmt.Println(quanteVolteRicorsiva(s,'z'))
}
