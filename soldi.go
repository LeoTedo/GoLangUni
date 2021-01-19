package main

import "fmt"

func main() {

	var valuta1 string
	var valuta2 string
	var tasso float32
	var cash float32
	fmt.Print("Inserisci la prima valuta:")
	fmt.Scan(&valuta1)
	fmt.Print("Inserisci la seconda valuta:")
	fmt.Scan(&valuta2)
	fmt.Print("Inserisci il tasso tra le due:")
	fmt.Scan(&tasso)
	fmt.Print("Inserisci il cash:")
	fmt.Scan(&cash)
	ris := cash * tasso
	fmt.Println(cash, valuta1, "Equivalgono a", ris, valuta2)
}
