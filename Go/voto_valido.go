package main
import "fmt"

/*Scrivere un programma Go voto_valido.go che legge un numero intero. Se il numero
non è compreso tra 0 e 30, stampa “voto non valido”, altrimenti non stampa niente.*/

func main()  {
  var voto int

  fmt.Print("voto: ")
  fmt.Scan(&voto)

  if voto >= 0 && voto <= 30{

  }else{
    fmt.Println("voto non valido")
  }
}
