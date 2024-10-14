package main
import "fmt"

/*Scrivere un programma trova.go che legge da standard input un char (rune)
e una stringa e stampa la posizione del char nella stringa,
o -1 se il char non c'*/

func main()  {
  var r rune
  var s string
  var flag = -1

  fmt.Scanf("%c",&r)
  //fmt.Println(r)
  fmt.Scan(&s)

  for i := 0; i < len(s); i++{
    if r == rune(s[i]){
      flag = i
      break
    }
  }

  fmt.Println(flag)
}
