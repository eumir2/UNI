package main

import "fmt"

func punti(s string) (punti int)  {
  punti = 0
  for i := 0 ; i < len(s); i++{
    switch string(s[i]) {
    case "A":
      punti += 11
    case "3":
      punti += 10
    case "K":
      punti += 4
    case "Q":
      punti += 3
    case "J":
      punti += 2
    default :
      punti += 0
    }
  }
  return
}


func main()  {
  var s string

  fmt.Scan(&s)
  if len(s) == 3{
    fmt.Printf("mano: %s: %d punti",s,punti(s))
  }else{
    fmt.Println("Stringa inserita non valida")
  }


}
