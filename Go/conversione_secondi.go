package main
import "fmt"

/*Scrivere un programma Go conversione_secondi.go che converta un numero dato
di secondi (fornito dall’utente) in giorni, ore, minuti, secondi senza mai usare l’operazione di
sottrazione.
Esempio di esecuzione:
numero di secondi: 123456
g:h:m:s = 1:10:17:36*/

func main()  {

  const sm = 60 //secondi in un minuto

  var secondi int

  so := sm * sm //secondi all'ora
  sg := so * 24 //secondi al giorno

  fmt.Print("numero di secondi: ")
  fmt.Scan(&secondi)

  giorni := secondi / sg
  sr := secondi % sg

  ore := sr / so
  sr = sr % so

  minuti := sr / sm
  sr = sr % sm

  fmt.Println("g:h:m:s =", giorni, ":", ore, ":", minuti, ":", sr)

}
