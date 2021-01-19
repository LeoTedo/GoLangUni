package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Find(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	//apro file morse.txt
	raffle, err := os.OpenFile("estrazione.txt", os.O_WRONLY|os.O_CREATE, 0666)
	outputWriter := bufio.NewWriter(raffle)
	if err != nil {
		fmt.Println("Hai sgravato")
	}
	defer raffle.Close()

	rand.Seed(time.Now().UnixNano())

	numamici := os.Args[1]
	finale, _ := strconv.Atoi(numamici)
	fmt.Println("Siete in", finale, "amici!")
	var amici []int
	for s := 1; s < finale+1; s++ {
		amici = append(amici, s)
	}
	//fmt.Println(amici)

	var regali []int
	for i := 0; i < len(amici)+1; i++ {
		if len(regali) == len(amici) {
			break
		} else {
			regalo := randInt(1, finale+1)
			_, trovato := Find(regali, regalo)
			if !trovato && regalo != amici[i] {
				regali = append(regali, regalo)
				//estr1 := "ciao"
				estr1 := strconv.Itoa(amici[i]) + "->" + strconv.Itoa(regalo) + "\n"
				//fmt.Println(amici[i], "->", regalo)
				//fmt.Println(estr1)
				outputWriter.WriteString(estr1)
			} else {
				//fmt.Println("doppione")
				i--
			}
		}

	}
	outputWriter.Flush()
	fmt.Println("Risultati nel file!")
}
