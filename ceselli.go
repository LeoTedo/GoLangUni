package main

import (
	"fmt"
	"strconv"
	"strings"
)

type linea struct {
	x  int
	y  int
	x2 int
	y2 int
}

func main() {
	l1 := []int{1, 2, 3, 4, 5, 765, 2, 87, 3}
	l2 := []int{5, 8, 12, 99, 666, 765, 3245, 2, 3}
	l3 := []int{1, 2, 3}
	l4 := []int{1, 2, 3, 2}
	c1 := linea{1, 2, -3, 2}
	c2 := linea{0, 6, 0, -6}
	//var ris []int
	fmt.Println(Primo(100))
	fmt.Println(Rev("ciao a tutti come va"))
	fmt.Println(TwoStrings("ciao", "canebe3qwyufdbweoufyeg"))
	fmt.Println(DoppiPari(l1, l2))
	fmt.Println(Repeat("vamo raga", 10))
	fmt.Println(Corrispondenza(l1, l2))
	fmt.Println(AnalisiSlice(l2))
	fmt.Println(Incrocio(c1, c2))
	fmt.Println(StringaInStringa("vamo raga ti aspecto", "vamo raga"))
	fmt.Println(Ordine(l3, l4))
	TrovaInt(l1, 3)
	fmt.Println(MagiaStringhe("ti aspecto", "vamo raga", "forza giuve"))
	fmt.Println(Crescente(l4))
	fmt.Println(Pari(l2))
	fmt.Println(PariR(l2, 0))
}

func Primo(n int) bool {
	divisori := 0
	for i := 1; i <= n; i++ {

		if n%i == 0 {
			divisori++
		}
	}
	if divisori == 2 {
		return true
	} else {
		return false
	}
}

func Rev(s string) string {
	sr := []rune(s)
	for i, j := 0, len(sr)-1; i < j; i, j = i+1, j-1 {
		sr[i], sr[j] = sr[j], sr[i]
	}
	return string(sr)
}

func TwoStrings(s1, s2 string) string {
	sr1, sr2 := []rune(s1), []rune(s2)

	for i := 0; i < len(s2); i++ {
		if strings.ContainsRune(s1, sr2[i]) == false {
			sr1 = append(sr1, sr2[i])
		}
	}
	return string(sr1)
}

func DoppiPari(l1, l2 []int) []int {
	var ris []int
	for i := 0; i < len(l1)-1; i++ {
		if l1[i]%2 == 0 {
			ris = append(ris, l1[i])
		}
	}
	for k := 0; k < len(l2)-1; k++ {
		if l2[k]%2 == 0 {
			ris = append(ris, l2[k])
		}
	}
	return ris
}

func Repeat(s string, n int) string {
	var ris string
	for i := 0; i < n; i++ {
		ris += s
	}
	return ris
}

func Corrispondenza(b1, b2 []int) []int {
	var ris []int
	mappa := make(map[int]bool)
	for i := 0; i < len(b1); i++ {
		for k := 0; k < len(b2); k++ {
			if b1[i] == b2[k] {
				mappa[b1[i]] = true
			}
		}
	}
	for i, _ := range mappa {
		ris = append(ris, i)
	}
	return ris
}

func AnalisiSlice(s []int) ([]int, []int, []int, []int) {
	var pos, neg, pari, dispari []int
	for i := 0; i < len(s); i++ {
		if s[i] > 0 {
			pos = append(pos, s[i])
		}
		if s[i] < 0 {
			neg = append(neg, s[i])
		}
		if s[i]%2 == 0 {
			pari = append(pari, s[i])
		}
		if s[i]%2 != 0 {
			dispari = append(dispari, s[i])
		}
	}
	return pos, neg, pari, dispari
}

func Incrocio(c1, c2 linea) bool {
	var trovato int
	if c1.x2 < c2.x && c2.x < c1.x {
		trovato++
	}
	if c2.y > c1.y && c1.y > c2.y2 {
		trovato++
	}
	fmt.Println(trovato)
	if trovato != 0 {
		return true
	} else {
		return false
	}
}

func abs(n int) int {
	if n < 0 {
		n = n * -1
	}
	return n
}

func StringaInStringa(s1, s2 string) bool {
	if strings.Contains(s1, s2) {
		return true
	} else {
		return false
	}
}

func Ordine(l1, l2 []int) bool {
	var stringa1, stringa2 string
	for i := 0; i < len(l1); i++ {
		conv := strconv.Itoa(l1[i])
		stringa1 += conv
	}
	for i := 0; i < len(l2); i++ {
		conv := strconv.Itoa(l2[i])
		stringa2 += conv
	}
	if stringa1 == stringa2 {
		return true
	} else {
		return false
	}
}

func TrovaInt(lista []int, n int) {
	for i := 0; i < len(lista); i++ {
		if lista[i] == n {
			fmt.Println("numero", n, "trovato in posizione", i)
		}
	}
}
func MagiaStringhe(s1, s2, s3 string) string {

	if strings.Contains(s1, s2) {
		return s1
	} else {
		return s1 + " " + s3
	}
}

func Crescente(l []int) bool {
	var check int
	for i := 0; i < len(l)-1; i++ {
		if l[i] < l[i+1] {
			check++
		}
	}
	//fmt.Println(check)
	if check == len(l)-1 {
		return true
	} else {
		return false
	}
}

func Pari(s []int) []int {
	var ris []int
	for i := 0; i < len(s); i++ {
		if s[i]%2 == 0 {
			ris = append(ris, s[i])
		}
	}
	return ris

}

func PariR(s []int) []int {
	var ris []int
	if len(s) == 0 {
		return 0
	} else {
		if PariR(s, i)[1:]%2 == 0 {
			ris = append(ris, PariR(s, i)[1:])
		}

	}

	i++
	return ris
}
