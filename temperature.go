package main

import (
	"fmt"
	"sort"
)

var temp int
var valori []int

func calcolamedia() float32 {
	totale := 0
	for i := 0; i < len(valori); i++ {
		totale += valori[i]
	}
	//fmt.Println("Totale:", totale)
	var media float32
	media = float32(totale / len(valori))
	return media
}

func calcolamediana() float32 {
	sort.Ints(valori)
	mediana := len(valori) / 2
	return float32(valori[mediana])
}

func valsotmedia() int {
	var sotto int
	for i := 0; i < len(valori); i++ {
		if valori[i] < int(calcolamedia()) {
			sotto++
		}
	}
	return sotto
}

func trebassi() []int {
	sort.Ints(valori)
	var trebassi []int
	for i := 0; i < 3; i++ {
		trebassi = append(trebassi, valori[i])
	}
	return trebassi
}

func trealti() []int {
	sort.Ints(valori)
	var trealti []int
	for i := len(valori) - 1; i >= len(valori)-3; i-- {
		trealti = append(trealti, valori[i])
	}
	return trealti
}

func main() {

	for {
		fmt.Println("inserisci la temperatura")
		fmt.Scanln(&temp)
		if temp != 999 {
			valori = append(valori, temp)
		} else {
			//fmt.Println(valori)
			fmt.Println("media delle temperature:", calcolamedia())
			fmt.Println("mediana delle temperature:", calcolamediana())
			fmt.Println("numero valori sotto la media:", valsotmedia())
			if len(valori) > 3 {
				fmt.Println("più di 3 valori!")
				fmt.Println("3 valori più bassi:", trebassi())
				fmt.Println("3 valori più alti:", trealti())
			} else {
				break
			}
			break
		}

	}

}
