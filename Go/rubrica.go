package main

import "fmt"

func main() {
	rubrica:=NuovaRubrica()

	rubrica=InserisciContatto(rubrica,"Rossi","Mario","Roma",1,"000",Milano,"000232")
	fmt.Println(rubrica)
}

type Indirizzo struct {
	via, cap, città string
	numeroCivico    uint
}
type Contatto struct {
	cognome, nome, telefono string
	indirizzo Indirizzo
}
type Rubrica []Contatto{

}
func NuovaRubrica()Rubrica{
	return Rubrica{}
}
func InserisciContatto(r Rubrica, cognome, nome string, via string, numero uint, cap, città
	string, telefono string) {
		return append (r,Contatto(cognome,nome,telefono,Indirizzo(via,cap,città,numero)))
}
