package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Comando struct {
	direzione string
	passi     int
}

func main() {
	invertire := LeggiComandi()
	bella := AnalizzaComandi(invertire)
	fmt.Println(bella)
	fmt.Println("Direzione più discorsa", Max(bella))
	Contrario(invertire)
}

func LeggiComandi() []Comando {
	var passaggi []Comando
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		var passo1 Comando
		passo1.direzione = scanner.Text()
		if scanner.Text() == "SUD" || scanner.Text() == "NORD" || scanner.Text() == "OVEST" || scanner.Text() == "EST" {
			scanner.Scan()
			numpasso, _ := strconv.Atoi(scanner.Text())
			passo1.passi = numpasso
			passaggi = append(passaggi, passo1)
		}

	}

	//fmt.Println(passaggi)
	return passaggi
}

func AnalizzaComandi(passaggi []Comando) map[string]int {
	ris := make(map[string]int)
	//var massimo []int
	for i := 0; i < len(passaggi); i++ {
		ris[passaggi[i].direzione] += passaggi[i].passi
	}
	return ris
}

func Max(m map[string]int) string {
	var massimoS string
	var massimo int
	for i, v := range m {
		if v > massimo {
			massimoS = i
			massimo = v
		} else {
			continue
		}
	}
	return massimoS
}

func Contrario(passaggi []Comando) {
	fmt.Println("Percorso al contrario:")
	for i := len(passaggi) - 1; i >= 0; i-- {
		if passaggi[i].direzione == "EST" {
			passaggi[i].direzione = "OVEST"

			fmt.Print(passaggi[i])
			continue
		}
		if passaggi[i].direzione == "OVEST" {
			passaggi[i].direzione = "EST"

			fmt.Print(passaggi[i])
			continue
		}
		if passaggi[i].direzione == "NORD" {
			passaggi[i].direzione = "SUD"
			fmt.Print(passaggi[i])
			continue
		}
		if passaggi[i].direzione == "SUD" {
			passaggi[i].direzione = "NORD"
			fmt.Print(passaggi[i])
			continue
		}
	}
	return
}
