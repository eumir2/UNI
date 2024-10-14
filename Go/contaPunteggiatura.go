package main

import (
  "fmt"
  "os"
  "bufio"
)

/*Scrivere un programma (contaPunteggiatura.go) che, dato un testo in input,
conta quante volte ciascuno dei seguenti caratteri di punteggiatura
(. , ; : ! ? " ') viene ripetuto.
Leggere un byte alla volta, utilizzando il metodo Split con
la funzione di split bufio.ScanBytes*/

func main()  {
  var m map [string]int
  m = make(map[string]int)

  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanBytes)

  for scanner.Scan(){
    l := scanner.Text()
    if l == "." || l == "," || l == ";" || l == ":" || l == "!" || l == "?" || l == "\"" || l == "'"{
      m[l]++
    }
  }


  fmt.Println(m)




}
