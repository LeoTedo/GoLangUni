package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	testo := LeggiTesto()
	mappaOccorrenze := Occorrenze(testo)
	fmt.Println("Istogramma: ")
	for carattere, occ := range mappaOccorrenze {
		StampaBarra(carattere, occ)
	}
}

func LeggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text()
	}
	return
}

func Occorrenze(s string) map[rune]int {
	occorrenze := make(map[rune]int)
	for _, r := range s {
		if unicode.IsLetter(r) {
			occorrenze[r]++
		}
	}
	return occorrenze
}

func StampaBarra(carattere rune, occ int) {
	barra := ""
	for i := 0; i < occ; i++ {
		barra += "*"
	}
	fmt.Printf("%c: %s\n", carattere, barra)
}
