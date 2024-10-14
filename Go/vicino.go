package main
import "fmt"

/*Scrivere un programma vicino.go che, data una serie di 5 valori
interi tra 0 e 20, trova il valore piu' vicino a TARGET, dove TARGET e`
una costante definita nel programma.
*/

func main()  {
  const TARGET = 10

  var vicino int //varibale dove salvo il numero più vicino
  var d_tmp int //variabile dove salvo la differenza temporanea
  var d_min int //variabile in cui salvo la differenza minima

  //input
  for i := 0; i < 5; i++{
    tmp := 0
    fmt.Print("valore ",i+1,": ")
    fmt.Scan(&tmp)


    if i == 0{
      if TARGET - tmp < 0 {
        d_min = (TARGET - tmp) * -1
      }else{
        d_min = TARGET - tmp
      }
    }else{
      if TARGET - tmp < 0 {
        d_tmp = (TARGET - tmp) * -1
      }else{
        d_tmp = TARGET - tmp
      }
      if d_tmp < d_min {
        d_min = d_tmp
        vicino = tmp
      }
    }
  }

  fmt.Println("Il valore più vicino a",TARGET,"è :", vicino)
}
