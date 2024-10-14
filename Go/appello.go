package main

import (
  "fmt"
  "os"
  "sort"
)

/*Scrivere un programma appello.go che legge da linea di comando
una sequenza di nomi e li stampa in ordine alfabetico.*/

func main()  {

//leggo i nomi da linea di comando e li metto in una slice
  s := os.Args[1:]

//ordino la slice
  sort.Strings(s)

//stampo la slice
  fmt.Println(s)

}
