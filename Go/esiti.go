package main
import "fmt"

/*Specifiche: Scrivere un programma esiti.go che associa voti in lettere a punteggi secondo la
seguente tabella. I voti sono int.
90 - 100 A
80 - 89 B
70 - 79 C
60 - 69 D
0 - 59 F*/

func main()  {
  var voto int
  var grado = ""

  fmt.Print("Voto: ")
  fmt.Scan(&voto)

  if voto < 0 || voto > 100{
    fmt.Println("voto non valido")
  }else if voto >= 0 && voto <= 59{
    grado = "F"
  }else if voto > 59 && voto <= 69{
    grado = "D"
  }else if voto > 69 && voto <= 79{
    grado = "C"
  }else if voto > 79 && voto <= 89{
    grado = "B"
  }else{
    grado = "A"
  }

fmt.Println(grado)

}
