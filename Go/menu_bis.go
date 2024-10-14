package main

import ("fmt"
        "os"
        "bufio")

func main() {

	const MENU = "Menu del giorno:\n" +
		"a. pizza\n" +
		"b. penne al pomodoro\n" +
		"c. cotoletta e patatine\n" +
		"d. crostata e caffe`" +
		"f. fine"

  var ordinazione = ""
	fmt.Println(MENU)

  fmt.Println("ordinazione?")

  scanner := bufio.NewScanner(os.Stdin)
  myLoop : for scanner.Scan(){
              scelta := scanner.Text()
              switch scelta {
            	case "a":
            		ordinazione += "\n pizza"
            	case "b":
            		ordinazione += "\n penne al pomodoro "
            	case "c":
            		ordinazione += "\n cotoletta e patatine "
            	case "d":
            		ordinazione += "\n crostata e caffe "
              case "f" :
                break myLoop
              default :
                fmt.Println("ordinazione non valida")
            	}
            }

	fmt.Print("hai ordinato", ordinazione)
}
