package main

import (
	"bufio"
	"fmt"
	"math"
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
	listaT := GeneraTriangoli(LeggiPunti())
	//fmt.Println(listaT)
	Verifica(listaT)
	//fmt.Println(best)
	//StringTriangolo(best)

}

func LeggiPunti() (punti []punto) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("inserisca!")
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		input := scanner.Text()
		res := strings.Split(input, ";")
		c1, _ := strconv.ParseFloat(res[1], 4)
		c2, _ := strconv.ParseFloat(res[2], 4)
		p := punto{res[0], c1, c2}
		punti = append(punti, p)
	}
	return punti
}

func GeneraTriangoli(punti []punto) (triangoli []triangolo) {

	if len(punti) == 3 {
		t := triangolo{punti[0], punti[1], punti[2]}
		triangoli = append(triangoli, t)
	} else {
		for i := 0; i < len(punti)-2; i++ {

			for k := i + 1; k < len(punti)-1; k++ {

				for x := k + 1; x < len(punti); x++ {
					t := triangolo{punti[i], punti[k], punti[x]}
					triangoli = append(triangoli, t)
				}
			}

		}
	}

	return triangoli
}

func Verifica(triang []triangolo) {
	var max float64
	var migliore triangolo
	for i := 0; i < len(triang); i++ {
		if triang[i].v1.p1 == triang[i].v2.p1 || triang[i].v2.p1 == triang[i].v3.p1 || triang[i].v1.p1 == triang[i].v3.p1 || triang[i].v1.p2 == triang[i].v2.p2 || triang[i].v2.p2 == triang[i].v3.p2 || triang[i].v1.p2 == triang[i].v3.p2 {
			if (triang[i].v1.p1 < 0 || triang[i].v2.p1 < 0 || triang[i].v2.p1 < 0 || triang[i].v3.p1 < 0 || triang[i].v1.p1 < 0 || triang[i].v3.p1 < 0 || triang[i].v1.p2 < 0 || triang[i].v2.p2 < 0 || triang[i].v2.p2 < 0 || triang[i].v3.p2 < 0 || triang[i].v1.p2 < 0 || triang[i].v3.p2 < 0) && (triang[i].v1.p1 > 0 || triang[i].v2.p1 > 0 || triang[i].v2.p1 > 0 || triang[i].v3.p1 > 0 || triang[i].v1.p1 > 0 || triang[i].v3.p1 > 0 || triang[i].v1.p2 > 0 || triang[i].v2.p2 > 0 || triang[i].v2.p2 > 0 || triang[i].v3.p2 > 0 || triang[i].v1.p2 > 0 || triang[i].v3.p2 > 0) {
				area := (Distanza(triang[i].v1, triang[i].v2) + Distanza(triang[i].v2, triang[i].v3)) / 2
				if area > max {
					migliore = triang[i]
				}
			}
		}
	}
	StringTriangolo(migliore)
}

func Distanza(p1, p2 punto) float64 {
	return math.Pow((math.Pow((p1.p1-p2.p1), 2) + math.Pow((p1.p2-p2.p2), 2)), 0.5)
}

func StringPunto(p punto) {
	fmt.Sprintln(p.nome, "= (", p.p1, ",", p.p2, ")")
}

func StringTriangolo(t triangolo) {
	var area float64
	area = (Distanza(t.v1, t.v2) + Distanza(t.v2, t.v3)) / 2
	if area == 0 {
		fmt.Print("Suca")
	} else {
		fmt.Println("Triangolo rettangolo con vertici", t.v1.nome, "= (", t.v1.p1, ",", t.v1.p2, "), ", t.v2.nome, "= (", t.v2.p1, ",", t.v2.p2, ") e", t.v3.nome, "= (", t.v3.p1, ",", t.v3.p2, ") ed area: ", area)

	}
}
