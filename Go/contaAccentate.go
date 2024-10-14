package main

import (
  "fmt"
  "os"
  "bufio"
)

/*Scrivere un programma (contaAccentate.go) che, dDato un testo in input,
conta quante lettere accentate ci sono (considerate solo à, è, é, ì, ò, ù).
Leggere una runa alla volta, utilizzando il metodo Split per Scanner
del package bufio con la funzione di split bufio.ScanRunes:
myScanner := bufio.NewScanner(os.Stdin)
myScanner.Split(bufio.ScanRunes)*/

func main()  {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanRunes)

  var lettereAccentate int

  for scanner.Scan(){
    l := scanner.Text()
    if l == "à" || l == "é" || l == "è" || l == "ì" || l == "ò" || l == "ù" {
      lettereAccentate++
    }
  }


  fmt.Println("Accentate trovate:",lettereAccentate)
}
