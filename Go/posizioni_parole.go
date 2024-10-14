package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)


/*
Scrivere un programma posizioni_parole.go
che legge una sequenza di stringhe da standard input e produce su
standard output, per ogni stringa, la lista delle posizioni in cui
essa compare nella sequenza (partendo dalla posizione 0)*/

//funzione che, data una mappa e una lista di parole aggiorna le posizioni di ciascuna parola
func aggiorna_posizioni(m map[string][]int,s []string)  {
  for i,v := range s {
    m[v] = append(m[v],i)
  }
}



func main()  {
  var posizioni map[string][]int
  var s_completa []string

  posizioni = make(map[string][]int)
  s_completa = make([]string,0)

  //leggo le righe
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
    l := scanner.Text()
    s := strings.Split(l, " ")
    s_completa = append(s_completa,s...)

  }

  //aggiorno le posizioni delle parole
    aggiorna_posizioni(posizioni,s_completa)

  //stampo la mappa
  fmt.Println(posizioni)



}
