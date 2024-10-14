package main
import "fmt"

func main()  {

  //variabili
  var n,i int
  fmt.Print("Quante volte:")
  fmt.Scan(&n)

  //controllo che il numero sia valido
  if n < 0 {
    fmt.Print("Numero inserito non valido")
  } else{
      for i = 0; i < n; i++{
      fmt.Print("Ciao\n" )
   }
  }
}
