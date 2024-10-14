package main
import "fmt"

/*Scrivere un programma tariffe.go che chiede all’utente l’età (int) e se è studente (bool)
e stampa il costo del biglietto di ingresso secondo la tabella che segue.
Utilizzare solo 4 (e non di più) istruzioni di assegnamento a una variabile
tariffa per poter poi stampare il costo del biglietto
età tariffa
[0 - 9) anni gratis
[9 - 18) anni 5
[18 - 26) anni 7.5
[26 - 65) anni 10
>= 65 anni 7.5
studenti >= 18 5*/

func main()  {
  var eta int
  var studente bool
  var tariffa float64


  fmt.Print("Età: ")
  fmt.Scan(&eta)

  fmt.Print("studente? (t/f): ")
  fmt.Scan(&studente)

  if (studente == true  && eta >= 18) || (eta >= 9 && eta < 18){
    tariffa = 5
  }else if (eta >= 65) || (eta >= 18 && eta < 26){
    tariffa = 7.5
  }else if eta >= 26 && eta < 65{
    tariffa = 10
  }else if eta >= 0 && tariffa < 9{
    tariffa = 0
  }

  if tariffa == 0{
    fmt.Println("gratis")
  }else{
    fmt.Println("ingresso", tariffa, "euro")
  }
}
