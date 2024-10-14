package main
import "fmt"

/*Scrivere una terza versione data_valida.go che legge tre interi gg, mm e aaaa (anno) e
verifica che sia una data valida.
 Questa versione finale tiene conto anche degli anni bisestili.*/


//funzione che preso l'anno ritorna true se è bisestile
func leapYear(anno int)bool  {
  return anno%400 == 0 || anno%4 == 0 && anno%100 != 0
}


func main()  {

  var gg, mm, aaaa int
  var lunghezza_mese int
  fmt.Print("gg mm aaaa? ")
  fmt.Scan(&gg, &mm, &aaaa)

  if mm > 0 && mm <= 12 && gg > 0 {
    if mm == 1 ||mm == 3 ||mm == 5 ||mm == 7 ||mm == 8 ||mm == 10 ||mm == 12{
      lunghezza_mese = 31
    }else if mm == 4 || mm == 6 || mm == 9 || mm == 11{
      lunghezza_mese = 30

    }else{
      if leapYear(aaaa){
        lunghezza_mese  = 29
      }else{
        lunghezza_mese = 28
      }
    }
  }


  if gg <= lunghezza_mese {
    fmt.Println("data valida")
  }else{
    fmt.Println("data non valida")
  }


}
