package main

import (
	"fmt"
	"strconv"
)

func main() {

	mappa := make(map[int]string)
	var num int
	var vocabolario string
	fmt.Scan(&num)
	fmt.Println("Inserisci il numero:", num)
	ns := strconv.Itoa(num)

	for i := 0; i < len(ns); i++ {
		ns2, _ := strconv.Atoi(string(ns[i]))
		if _, ok := mappa[ns2]; ok {
			continue
		} else {
			fmt.Println("Inserisci il numero in caratteri:")
			fmt.Scan(&vocabolario)
			mappa[ns2] = vocabolario
		}

	}

	for i := 0; i < len(ns); i++ {
		ns2, _ := strconv.Atoi(string(ns[i]))
		if val, ok := mappa[ns2]; ok {
			fmt.Print(val, "-")
		}
	}
}
