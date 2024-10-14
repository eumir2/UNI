package main
import "fmt"


//programma che dato l'anno stabilisce se una persona è maggiorenne

func main()  {
  //variabili
  var annoN int
  var anno = 2021

  fmt.Print("Inserire l'anno di nascita: ")
  fmt.Scan(&annoN)

  if anno - annoN >= 18 {
    fmt.Println("Maggiorenne")
  }else{
    fmt.Println("Minorenne")
  }

}
