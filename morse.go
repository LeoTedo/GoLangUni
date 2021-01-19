package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	var parola string
	var ris []string
	for {
		//apro file morse.txt
		morse, err := os.Open("morse.txt")
		//dichiaro ed inizializzo il reader sul file
		input := bufio.NewReader(morse)
		if err != nil {
			fmt.Println("Hai sgravato")
		}
		defer morse.Close()
		fmt.Println("inserisci parola")
		fmt.Scan(&parola)
		if parola == "stop" {
			break
		} else {
			parolaR := []rune(parola)
			fmt.Println("la tua parola", string(parolaR))
			for i := 0; i < len(parolaR); i++ {
				for {
					rigacodice, rerr := input.ReadString('\n')
					morseR := []rune(rigacodice)
					if len(morseR) > 0 {
						if parolaR[i] == morseR[0] {
							ris = append(ris, rigacodice)
							fmt.Print(string(morseR[3:len(morseR)-2]), "/")
							break
						}
					}
					if rerr == io.EOF {
						break
					}
				}
				morse.Seek(0, io.SeekStart)
			}
		}
		morse.Close()
	}
}
