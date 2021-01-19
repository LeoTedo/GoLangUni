package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	lista := os.Args[0:]
	for i := 0; i < len(lista); i++ {
		fmt.Print(TrasformaParola(lista[i], i), " ")
	}
}

func TrasformaParola(parola string, posizione int) (parolaTrasformata string) {
	parolaE := []rune(parola)
	var ptr []rune
	if posizione%2 == 0 {
		for i := 0; i < len(parolaE); i++ {
			if i == 0 {
				ptr = append(ptr, unicode.ToUpper(parolaE[i]))
			} else {
				ptr = append(ptr, unicode.ToLower(parolaE[i]))
			}
		}
	} else {
		for i := 0; i < len(parolaE); i++ {
			if i%2 == 0 {
				ptr = append(ptr, unicode.ToLower(parolaE[i]))
			} else {
				ptr = append(ptr, unicode.ToUpper(parolaE[i]))
			}
		}
	}
	//fmt.Println(string([]rune(parolaE)))
	parolaTrasformata = string([]rune(ptr))
	return parolaTrasformata
}
