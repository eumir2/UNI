package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"text/scanner"
)

/* originale Anna
func main() {
	var scelta string
	var countLines int
	diz := make(map[string][]int)
	fine := false

	bfrd := bufio.NewReader(os.Stdin)
	printMenu()
	for !fine {
		fmt.Scan(&scelta)
		fmt.Println("scelta", scelta)
		switch scelta {
		case "+":
			riga, _ := bfrd.ReadString('\n')
			//fmt.Println("riga:", riga)
			countLines++
			parole := strings.Fields(riga)
			for _, wrd := range parole {
				diz[wrd] = append(diz[wrd], countLines)
			}
			//fmt.Println(diz)
		case "?":
			parola, _ := bfrd.ReadString('\n')
			parola = parola[:len(parola)-1]
			fmt.Println("parola:", parola)
			fmt.Println("righe", diz[parola])
		case "l":
			fmt.Println(countLines)
		case "n":
			fmt.Println(len(diz))
		case "p":
			fmt.Println(diz)
		case "f":
			fine = true
		default:
			fmt.Println("invalid option")
		}
	}
}
*/

func main() {
	//TODO implementare in modo da verificare test
	var countLines int
	diz := make(map[string][]int)
	//fine := false

	sc := bufio.NewScanner(os.Stdin)
	printMenu()
	for sc.Scan() {
		//fmt.Println("scelta", sc.Text())
		if len(sc.Text()) > 0 {
			switch sc.Text()[0] {
			case 'r':
				fmt.Println("riga:", sc.Text())
			case '+':
				parole := strings.Fields(sc.Text()[1:])
				countLines++
				for _, wrd := range parole {
					diz[wrd] = append(diz[wrd], countLines)
				}

				/*
						riga, _ := bfrd.ReadString('\n')
						//fmt.Println("riga:", riga)
						countLines++
						parole := strings.Fields(riga)
						for _, wrd := range parole {
							diz[wrd] = append(diz[wrd], countLines)
						}
						//fmt.Println(diz)
					case "?":
						parola, _ := bfrd.ReadString('\n')
						parola = parola[:len(parola)-1]
						fmt.Println("parola:", parola)
						fmt.Println("righe", diz[parola])
					case "l":
						fmt.Println(countLines)
					case "n":
						fmt.Println(len(diz))
					case "f":
						fine = true
				*/
			case '?':
				parola := sc.Text()[2:]
				//parola = parola[:len(parola)-1]
				fmt.Println("parola:", parola)
				fmt.Println("righe", diz[parola])
			case 'p':
				fmt.Println(diz)
			default:
				fmt.Println("invalid option")
			}
		}
	}
}

func printMenu() {
	// mette uno spazio di troppo?!?
	fmt.Println("+ (legge e memorizza)\n",
		"? (numeri di riga in cui è comparsa la parola data)\n",
		//"l (Stampa il numero di righe lette)\n",
		//"n (Stampa il numero di parole distinte lette)\n",
		"p (Stampa le parole)")
	//"f (Termina l’esecuzione)")
	return
}
