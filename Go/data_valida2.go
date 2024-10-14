package main
import "fmt"

/*Scrivere una seconda versione data_valida2.go che legge due interi gg e mm
rappresentanti giorno e mese dell’anno, e verifica che sia una data valida. In questa seconda
versione si tenga conto del fatto che solo i mesi 1, 3, 5, 7, 8, 10, 12 hanno 31 giorni, che i mesi 4, 6,
9, 11 ne hanno 30, e si assuma che febbraio ne abbia sempre 28. Il programma stampa "data
valida" / "data non valida".*/

func main()  {
  var gg, mm int
  var lunghezza_mese int
  fmt.Print("giorno e mese: ")
  fmt.Scan(&gg, &mm)

  if mm > 0 && mm <= 12 && gg > 0 {
    if mm == 1 ||mm == 3 ||mm == 5 ||mm == 7 ||mm == 8 ||mm == 10 ||mm == 12{
      lunghezza_mese = 31
    }else if mm == 4 || mm == 6 || mm == 9 || mm == 11{
      lunghezza_mese = 30

    }else{
      lunghezza_mese = 28
    }
  }
  if gg <= lunghezza_mese {
    fmt.Println("data valida")
  }else{
    fmt.Println("data non valida")
  }


}
