package main

import (
  "fmt"
  "os"
  "bufio"
  "unicode/utf8"
)

/*Scrivere un programma (lungPariDispari.go) che, dato un testo in input,
stampa prima tutte le righe che contengono un numero pari di rune (rune, non byte)
e poi tutte quelle che ne contengono un numero dispari.*/

func main()  {
  scanner := bufio.NewScanner(os.Stdin)

  var rPari, rDispari []string

  runeInRiga := 0
  for scanner.Scan(){
    l := scanner.Text()

    //conto le rune in ciascuna riga
    runeInRiga = utf8.RuneCountInString(l)
    fmt.Println(runeInRiga)
    //aggiungo la riga alla rispettiva slice
    if runeInRiga % 2 == 0{
      rPari = append(rPari,l)
    }else{
      rDispari = append(rDispari, l)
    }
    runeInRiga = 0
  }

  fmt.Println("Linee di lunghezza pari:\n",rPari)
  fmt.Println("--------------------------------")
  fmt.Println("Linee di lunghezza dispari:\n",rDispari)
}
