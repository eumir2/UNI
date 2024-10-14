package main

import ("fmt")

/*Scrivi un programma stampa.go che legge una stringa da standard input e la stampa in verticale sia dall'inizio alla fine che al contrario.
 Caricalo poi su upload.*/

 func main()  {
   var s string
   var sRev string

   fmt.Scan(&s)

    j := len(s)-1

   for i:= 0 ; i < len(s); i++ {
     fmt.Println(string(s[i]))
     sRev += string(s[j]) + "\n"
     j--
   }
   fmt.Println(sRev)
 }
