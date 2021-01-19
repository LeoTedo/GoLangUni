package main

import "fmt"

func main() {

	var a string
	a = "ciao"
	var b *string
	b = &a
	fmt.Println(a, *b)

}
