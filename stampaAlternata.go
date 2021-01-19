package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	const testo = `cominciò a gridar la fiera bocca, 
	E ’l duca mio ver lui: «Anima sciocca, 
	quand’ira o altra passion ti tocca!
	Raphél maì amèche zabì almi, 
	cui non si convenia più dolci salmi.  
	tienti col corno, e con quel ti disfoga`
	var libro []string
	scanner := bufio.NewScanner(strings.NewReader(testo))
	//scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		libro = append(libro, scanner.Text())
		//fmt.Println(scanner.Text(), i)
	}
	fmt.Println("Righe pari:", "\n")
	for i := 1; i < len(libro)+1; i++ {
		if i%2 == 0 {
			fmt.Println(libro[i-1])
		}
	}

	fmt.Println("Righe dispari:", "\n")
	for i := 0; i < len(libro); i++ {
		if i%2 == 0 {
			fmt.Println(libro[i])
		}

	}

}
