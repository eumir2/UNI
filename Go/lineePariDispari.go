package main

import (
  "fmt"
  "os"
  "bufio"
)

/*Scrivere un programma (lineePariDispari.go) che, dato un testo in input,
stampa prima tutte le righe pari e poi tutte le dispari.
La prima riga del testo è la riga 1.*/

func main()  {
  var rPari, rDispari []string
  var numRiga int
  numRiga  = 1

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
    l := scanner.Text()
    if numRiga % 2 == 0{
      rPari = append(rPari, l)
    }else{
      rDispari = append(rDispari, l)
    }
    numRiga++
  }

  //stampo le righe pari
  fmt.Println(rPari)

  //stampo le righe dispari
  fmt.Println(rDispari)
}
