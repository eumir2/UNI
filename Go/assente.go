package main

import "fmt"
import "os"

/*che, dato come argomenti una stringa e un byte,
 restituisce TRUE se la stringa non contiene il byte e FALSE altrimenti;*/
func assente(s string, b byte) bool{
  //caso base
  if len(s) == 0{
    return false
  }

  //ricorsione
  if s[0] == b{
    return true
  }else{
    return assente(s[1:],b)
  }
}

func main()  {
  s := os.Args[1]
  bString := os.Args[2]

  b := []byte(bString)

  //fmt.Println(s, b)

  fmt.Println(assente(s,b[0]))
}
