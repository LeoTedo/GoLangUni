package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	testo := LeggiTesto()
	m := (Occorrenze(testo))

	for k, v := range m {
		fmt.Print(string(k), ":")
		for i := 0; i < v; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

func LeggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + " "
	}
	return
}

func Occorrenze(testo string) (m map[rune]int) {
	m = make(map[rune]int)
	for _, v := range testo {
		if unicode.IsLetter(v) {
			m[v]++
		}
	}
	return
}
