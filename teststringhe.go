package main

import (
	"fmt"
	"strings"
)

func main() {

	parola := "Giuseppe"

	if strings.ContainsRune(parola, 'x') {
		fmt.Println("Vaffanculo Giuseppe")
	} else {
		fmt.Println("Nooooooo Giuseppe")
	}

}
