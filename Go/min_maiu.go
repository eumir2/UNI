package main
import "fmt"

/*Scrivere un programma minu_maiu.go che legge da standard input una stringa
e stampa se la stringa contiene solo minuscole, solo maiuscole o sia minuscole
che maiuscole.*/

func main()  {
  var s string
  var cM, cm int

  fmt.Scan(&s)

  for _,r := range s{
    if r >= 'a' && r <= 'z'{
      cm++
    }else if r >= 'A' && r <= 'Z'{
      cM++
    }
  }

  if cm == len(s){
    fmt.Println("minuscole")
  }else if cM == len(s){
    fmt.Println("maiuscole")
  }else{
    fmt.Println("mista")
  }
}
