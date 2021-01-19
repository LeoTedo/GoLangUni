package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	ris := NuovoTragitto()
	fmt.Println(ris)
	ris2 := lunghezza(ris)
	fmt.Println("lunghezza percorso:", ris2)
	fmt.Println(HalfString(HalfPunto(ris)))
}

type punto struct {
	nome string
	x    float64
	y    float64
}

var input string

func NuovoTragitto() (tragitto []punto) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("inserisca!")
		scanner.Scan()
		input = scanner.Text()
		//fmt.Println(input)
		res := strings.Split(input, ";")
		if input != "" {
			//fmt.Println(res[0], res[1], res[2])
			f, _ := strconv.ParseFloat(res[1], 4)
			f2, _ := strconv.ParseFloat(res[2], 4)
			p := punto{res[0], f, f2}
			tragitto = append(tragitto, p)
		} else {
			break
		}
	}

	return tragitto
}

func distanza(p1, p2 punto) float64 {
	return math.Pow((math.Pow((p1.x-p2.x), 2) + math.Pow((p1.y-p2.y), 2)), 0.5)
}

func lunghezza(tragitto []punto) float64 {
	var totale float64
	for i := 0; i < len(tragitto)-1; i++ {
		totale += distanza(tragitto[i], tragitto[i+1])
	}
	return totale
}

func HalfPunto(tragitto []punto) punto {
	return tragitto[(len(tragitto)/2)+1]
}

func HalfString(p punto) string {
	//fmt.Println("punto oltre metà:", p.nome, "=", "(", p.x, ",", p.y, ")")
	return fmt.Sprintln("punto oltre metà:", p.nome, "=", "(", p.x, ",", p.y, ")")
}
