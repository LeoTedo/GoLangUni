package main

import (
	"fmt"
	"os"
	"sort"
)

type ByLen []string

func (a ByLen) Len() int           { return len(a) }
func (a ByLen) Less(i, j int) bool { return len(a[i]) < len(a[j]) }
func (a ByLen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	lista := os.Args[1]
	parolaE := []rune(lista)
	occorrenze := 0
	//var dump string
	var ris []string
	for i := 0; i < len(parolaE)-1; i++ {
		for k := len(parolaE); k > i; k-- {
			//fmt.Println(string(parolaE[i:k]))
			if parolaE[i] == parolaE[k-1] && len(parolaE[i:k]) >= 3 {
				ris = append(ris, string(parolaE[i:k]))
				//occorrenze++
				//fmt.Print(string(parolaE[i]))
				//fmt.Println(string(parolaE[i:k]), "-> Occorrenze", occorrenze)
			}
		}

	}
	sort.Sort(ByLen(ris))
	//fmt.Println(ris)
	for i, j := 0, len(ris)-1; i < j; i, j = i+1, j-1 {
		ris[i], ris[j] = ris[j], ris[i]
	}
	for i := 0; i < len(ris); i++ {
		occorrenze = 0
		occorrenze++
		if ris[i] != ris[len(ris)-1] && ris[i] == ris[i+1] {
			occorrenze++
			i++
		}
		fmt.Println(string(ris[i]), "-> Occorrenze", occorrenze)
	}
}
