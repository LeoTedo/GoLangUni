package main

import (
	"fmt"
	"os"
	"sort"
)

type SortBy []string

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func main() {

	lista := os.Args[1:]

	sort.Sort(SortBy(lista))

	fmt.Println(lista)

}
