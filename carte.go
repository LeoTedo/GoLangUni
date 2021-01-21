package main

import (
	"fmt"
	"math/rand"
	"time"
)

type carta struct {
	valore string
	seme   string
}

var mazzo = []carta{
	{"A", "\u2665"}, {"1", "\u2665"}, {"2", "\u2665"}, {"3", "\u2665"}, {"4", "\u2665"}, {"5", "\u2665"}, {"6", "\u2665"}, {"7", "\u2665"}, {"8", "\u2665"}, {"9", "\u2665"}, {"10", "\u2665"}, {"J", "\u2665"}, {"Q", "\u2665"}, {"K", "\u2665"},
	{"A", "\u2666"}, {"1", "\u2666"}, {"2", "\u2666"}, {"3", "\u2666"}, {"4", "\u2666"}, {"5", "\u2666"}, {"6", "\u2666"}, {"7", "\u2666"}, {"8", "\u2666"}, {"9", "\u2666"}, {"10", "\u2666"}, {"J", "\u2666"}, {"Q", "\u2666"}, {"K", "\u2666"},
	{"A", "\u2663"}, {"1", "\u2663"}, {"2", "\u2663"}, {"3", "\u2663"}, {"4", "\u2663"}, {"5", "\u2663"}, {"6", "\u2663"}, {"7", "\u2663"}, {"8", "\u2663"}, {"9", "\u2663"}, {"10", "\u2663"}, {"J", "\u2663"}, {"Q", "\u2663"}, {"K", "\u2663"},
	{"A", "\u2660"}, {"1", "\u2660"}, {"2", "\u2660"}, {"3", "\u2660"}, {"4", "\u2660"}, {"5", "\u2660"}, {"6", "\u2660"}, {"7", "\u2660"}, {"8", "\u2660"}, {"9", "\u2660"}, {"10", "\u2660"}, {"J", "\u2660"}, {"Q", "\u2660"}, {"K", "\u2660"}}

const semi = 4
const valori = 13
const mazzoN = 52

var v int

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Dimmi un numero bro")
	fmt.Scan(&v)
	ScegliCarta(v)
	fmt.Println("Ora la carta estratta per te!", EstraiCarta())
	fmt.Println("La tua mano da poker", Dai4Carte())
}

func ScegliCarta(v int) {
	fmt.Println("La tua carta è:", mazzo[v])
	return
}

func EstraiCarta() carta {
	estrazione := rand.Intn(mazzoN)
	return mazzo[estrazione]
}

func Dai4Carte() []carta {
	var poker []carta
	for i := 0; i < 4; i++ {
		poker = append(poker, EstraiCarta())
	}
	return poker
}
