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
	tienti col corno, e con quel ti disfoga?????`
	//accenti := 0
	scanner := bufio.NewScanner(strings.NewReader(testo))
	scanner.Split(bufio.ScanBytes)
	/*for scanner.Scan() {
		if scanner.Text() == "à" || scanner.Text() == "è" || scanner.Text() == "é" || scanner.Text() == "ì" || scanner.Text() == "ò" || scanner.Text() == "ù" {
			accenti++
		}
	}*/
	m := make(map[string]int)
	for _, v := range testo {
		if v == '.' || v == ',' || v == ';' || v == ':' || v == '!' || v == '?' || v == '"' {
			sv := string(v)
			m[sv]++
			//fmt.Println(string(v))
		}
	}

	fmt.Println(m)
}
