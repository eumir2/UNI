package main

import ("fmt"
        "bufio"
        "os")

/*Scrivere un programma vocali.go che analizza un testo e conta
le occorrenze delle vocali (sia minuscole che maiuscole,
ma non le accentate) nel testo e stampa, ma solo per le vocali
presenti nel testo, il numero di volte che le vocali stesse sono
presenti nel testo.

In particolare il programma è dotato di:

- una funzione
	func contaVocali(s string, vocali map[rune]int)
per contare le occorrenze delle diverse vocali
(sia minuscole che maiuscole - vedi es sotto)
in tutte le stringhe che le vengono passate.
La funzione, data una stringa s e una mappa vocali,
aggiorna opportunamente la mappa vocali aggiungendo
eventuali vocali non ancora presenti / incrementandone i valori.*/

func contaVocali(s string, vocali map[rune]int)  {
  for _,c := range s{
    switch c {
    case 'a', 'e', 'i','o','u','A', 'E','I','O','U' :
        vocali[c]++
    }
  }
}

func main()  {
  var vocali map[rune]int
  vocali = make(map[rune]int)


  //leggo le stringhe
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
    t := scanner.Text()
    contaVocali(t, vocali)
    for k,v := range vocali{
      fmt.Println(string(k), " : ",v)
    }
  }
}
