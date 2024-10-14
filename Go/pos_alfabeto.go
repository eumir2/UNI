package main
import "fmt"

/*Scrivere un programma pos_alfabeto.go che
- legge una sequenza di caratteri ASCII (byte) terminati da '.' (es: 5Nov.),
- per ciascuno stabilisce se è una lettera minuscola (es. f) e stampa
un messaggio per riga: se minuscola, il carattere e la sua posizione
nell'alfabeto (es "f è la 6^a"); altrimenti "altro"
- quando termina, stampa "bye"*/

func main()  {
  var carattere byte

  for {
    fmt.Scanf("%c", &carattere)
    if carattere == '.'{
      break
    }

    //identifico se è maiuscolo o minuscolo
    if carattere >= 'a' && carattere <= 'z'{
      fmt.Print(string(carattere), " è la ")
      fmt.Print(carattere - 'a' + 1)
      fmt.Println("^a")
    }else{
      fmt.Println("altro")
    }
  }
  fmt.Println("bye")
}
