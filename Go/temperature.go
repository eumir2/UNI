package main

import (
  "fmt"
  "sort"
)
/*Scrivere un programma temperature.go che legge delle temperature
(int) da tastiera e termina quando il valore letto è 999.
Il programma deve stampare
- la media
- la mediana
- il numero di temperature sotto la media delle temperature stesse.
- le tre temperature più basse (se ci sono almeno 3 temperature)
- le tre temperature più alte (se ci sono almeno 3 temperature)
*/

//funzione che data la slice, calcola la media dei valori
//restituisce 0,false se la slice è vuota
func media(t []int)(int,bool)  {
  tmp := 0
  l := len(t)

  if l == 0{
    return 0 , false
  }

  for i:= 0; i < l; i++{
    tmp += t[i]
  }
  return tmp / l , true
}


/*Nota 1: la mediana di un insieme di dati  e` data, nel
caso ci sia un numero dispari di dati, dal dato centrale
dei dati ordinati per valore (ad es. crescente), altrimenti
dalla media dei due dati centrali*/
//funzione che data una slice ordinata, calcola e restituisce la mediana
//restituisce 0,false se la slice è vuota
func mediana(t []int) (int,bool)  {
  l := len(t)
  middle := l / 2
  if l == 0{
    return 0 , false
  }

  if l % 2 != 0{
    return t[middle] , true
  }else{
    return (t[middle]+t[middle+1])/2 , true
  }
}

func main()  {
  var temperature []int
  temperature = make([]int, 0)

  //leggo i valori in input e popolo la slice
  var tmp int
  for{
    fmt.Scan(&tmp)
    if tmp == 999{
      break
    }
    temperature = append(temperature, tmp)
  }

  //ordino subito la slice in ordine crescente
  sort.Ints(temperature)

  //fmt.Println(temperature)

  //calcolo la media e la stampo
  media, ok := media(temperature)
  if ok{
    fmt.Println("media :",media)
  }else{
    fmt.Println("lista vuota")
  }

  //calcolo la mediana e la stampo
  mediana, ok := mediana(temperature)
  if ok{
    fmt.Println("mediana :",mediana)
  }else{
    fmt.Println("lista vuota")
  }

  //temperature piu basse della media
  fmt.Print("temperature sotto la media: ")
  for i := 0; temperature[i] < media; i++{
    fmt.Print(temperature[i], ", ")
  }
  fmt.Println()

  //tre temperature piu alte e piu basse
  if len(temperature) >= 3 {
    fmt.Println("tre temperature più basse:",temperature[:3])
    fmt.Println("tre temperature più alte:",temperature[len(temperature)-3:])
  }


}
