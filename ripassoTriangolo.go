package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type punto struct {
	nome string
	p1   float64
	p2   float64
}

type triangolo struct {
	v1 punto
	v2 punto
	v3 punto
}

func main() {
	fmt.Println(LeggiPunti())
}

func LeggiPunti() (punti []punto) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("inserisca!")
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			break
		} else {
			res := strings.Split(input, ";")
			c1, _ := strconv.ParseFloat(res[1], 4)
			c2, _ := strconv.ParseFloat(res[2], 4)
			p := punto{res[0], c1, c2}
			punti = append(punti, p)
		}

	}
	return punti
}

func GeneraTriangoli(punti []punto) (triangoli []triangolo) {

	for i := 0; i < len(punti); i++ {

	}

	return triangoli
}
