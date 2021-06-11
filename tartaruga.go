package main

import (
	"fmt"
	"strconv"
)

type turtle struct {
	x   int
	y   int
	dir byte
}

func main() {
	koopa := turtle{0, 0, 'N'}
	var comando string
	var ris []turtle
	for {
		fmt.Scan(&comando)
		if comando == "stop" {
			for i := 0; i < len(ris); i++ {
				fmt.Println("(", ris[i].x, ",", ris[i].y, ") dir", string(ris[i].dir))
			}
			//fmt.Println(ris)
			//fmt.Println(koopa)
			break
		} else {
			comando2, _ := strconv.Atoi(string(comando[2]))
			//direzione, _ := strconv.Atoi(string(comando[2]))
			//fmt.Println(comando2)
			//fmt.Println(string(comando[0]))
			switch string(comando[0]) {
			case "F":
				//koopa.y += comando2
				koopa = forward(koopa, comando2)
			case "B":
				koopa = backward(koopa, comando2)
			case "S":
				koopa = setDir(koopa, comando[2])
			}

			ris = append(ris, koopa)
		}
	}

}

func forward(koopa turtle, passi int) turtle {

	switch string(koopa.dir) {
	case "N":
		koopa.y += passi
	case "S":
		koopa.y -= passi
	case "E":
		koopa.x += passi
	case "O":
		koopa.x -= passi
	}
	return koopa
}

func backward(turtle turtle, passi int) turtle {
	switch string(turtle.dir) {
	case "N":
		turtle.y -= passi
	case "S":
		turtle.y += passi
	case "E":
		turtle.x -= passi
	case "O":
		turtle.x += passi
	}
	return turtle
}

func setDir(turtle turtle, dir byte) turtle {
	turtle.dir = dir
	return turtle
}
