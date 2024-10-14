package main

import (
  "fmt"
  "os"
  "bufio"
)
/*Scrivere un programma (parolaPresente.go) che verifica se una parola,
data come primo argomento da linea di comando, è presente
(come parola o anche come sottostringa di una parola) in un file di testo,
il cui nome è passato come secondo argomento da linea di comando.
Nel caso in cui lo sia, il programma stampa in che posizione si trova la parola
(o la sua sovrastringa) nel testo, altrimenti stampa "not found".
Partiamo a contare da 1, quindi la prima parola del testo si trova in posizione 1.
Nel caso in cui manchino (uno o tutti e due) parametri da linea di comando, il programma
stampa "args missing".
Suggerimento: usare Split(bufio.ScanWords)*/

func main()  {
  parola := os.Args[1]
  nomeFile := os.Args[2]

  //fmt.Println(parola)
  //fmt.Println(file)
  if len(parola) == 0 ||

  //apro il file e lo leggo
	file, err := os.Open(nomeFile)
	if err != nil {
		fmt.Println("file not found")
		return
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
  fileScanner.Split(bufio.ScanWords)

  posParola := 1
  for fileScanner.Scan(){
    l := fileScanner.Text()
    if parola == l{
      fmt.Println("found at:",posParola)
      break
    }else{
      posParola++
    }
  }
}
