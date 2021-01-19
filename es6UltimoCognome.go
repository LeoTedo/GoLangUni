package main

import "fmt"

func main() {

	var cognome string

	for {
		var ins string = ""
		fmt.Println("Inserisci cognome:")
		fmt.Scanln(&ins)
		if ins == "" {
			break
		} else {
			if ins > cognome {
				cognome = ins
			}
		}

	}
	fmt.Println("ultimo cognome: ", cognome)
}
