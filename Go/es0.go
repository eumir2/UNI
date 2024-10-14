package main
import "fmt"

/*Scrivere un programma es0.go che legge un byte e
- lo stampa (occorre Scanf in più per catturare l'invio)
- stampa precedente, byte stesso, e successivo in ordine
lessicografico (ASCII). Ad es. per 'd': cde
- stabilisce se è una lettera tra A e L, o altro (stampa "A-L" o "altro")
poi legge una stringa (di rune) e la stampa in verticale. Ad esempio città*/


func main()  {
  var carattere byte
  var s string

  fmt.Scan(&carattere)
  //stampo il precedente, il carattere e il successivo
  fmt.Printf("%c %c %c \n",carattere -1, carattere, carattere +1)

  //controllo che sia compreso tra A e L
  if carattere >= 65 && carattere <= 76{
    fmt.Println("A-L")
  }else{
    fmt.Println("altro")
  }

  //leggo una stringa e la stampo lettera per lettera
  fmt.Scan(&s)
  for i := 0; i < len(s); i++ {
    fmt.Printf("%c \n", s[i])
  }
}
