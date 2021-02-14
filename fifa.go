package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type calciatore struct {
	ovr  int
	pos  string
	nome string
	vel  int
	tir  int
	pas  int
	dri  int
	def  int
	phy  int
}

func main() {
	var database []calciatore
	var menu int
	scanner := bufio.NewScanner(os.Stdin)
	database = []calciatore{{88, "AT", "Depay", 89, 87, 88, 87, 35, 83}, {91, "CC", "Camavinga", 86, 84, 91, 90, 87, 87}, {93, "AD", "Messi", 85, 92, 91, 95, 38, 65}, {85, "AS", "Rashford", 91, 83, 78, 85, 45, 78}}

	fmt.Println("1. Cerca giocatore \n", "2. Confronta giocatori")
	fmt.Scanln(&menu)
	switch menu {
	case 1:
		fmt.Println("Inserisci il nome del giocatore che vuoi cercare")
		scanner.Scan()
		Search(scanner.Text(), database)
	case 2:
		var g1, g2 int
		fmt.Println("Inserisci i numeri corrispondenti ai giocatori da confrotare")
		for i, v := range database {
			fmt.Println(i, v.nome)
		}
		fmt.Scan(&g1, &g2)
		Confronto(database[g1], database[g2])
	case 3:
		break
	}

}

func Search(s string, d []calciatore) {
	var trovati int
	for _, v := range d {
		if strings.Contains(strings.ToLower(v.nome), strings.ToLower(s)) {
			trovati++
			fmt.Println("OVR:", v.ovr)
			fmt.Println("Ruolo:", v.pos)
			fmt.Println("Nome:", v.nome)
			fmt.Println("Vel:", v.vel)
			fmt.Println("Tir:", v.tir)
			fmt.Println("Pas:", v.pas)
			fmt.Println("Dri", v.dri)
			fmt.Println("Def", v.def)
			fmt.Println("Phy", v.phy)
		}
	}
	if trovati == 0 {
		fmt.Println("Nessun risultato trovato!")
	}
}

func Confronto(c1, c2 calciatore) {
	fmt.Println("Confronto tra", c1.nome, "e", c2.nome, "!")
	if c1.ovr > c2.ovr {
		fmt.Println("Meglio", c1.nome, "con", c1.ovr, "di overall")
	} else {
		fmt.Println("Meglio", c2.nome, "con", c2.ovr, "di overall")
	}
	if c1.vel > c2.vel {
		fmt.Println("Meglio", c1.nome, "con", c1.vel, "di velocità")
	} else {
		fmt.Println("Meglio", c2.nome, "con", c2.vel, "di velocità")
	}
	if c1.tir > c2.tir {
		fmt.Println("Meglio", c1.nome, "con", c1.tir, "di tiro")
	} else {
		fmt.Println("Meglio", c2.nome, "con", c2.tir, "di tiro")
	}
	if c1.pas > c2.pas {
		fmt.Println("Meglio", c1.nome, "con", c1.pas, "di passaggio")
	} else {
		fmt.Println("Meglio", c2.nome, "con", c2.pas, "di passaggio")
	}
	if c1.dri > c2.dri {
		fmt.Println("Meglio", c1.nome, "con", c1.dri, "di dribbling")
	} else {
		fmt.Println("Meglio", c2.nome, "con", c2.dri, "di dribbling")
	}
	if c1.def > c2.def {
		fmt.Println("Meglio", c1.nome, "con", c1.def, "di difesa")
	} else {
		fmt.Println("Meglio", c2.nome, "con", c2.def, "di difesa")
	}
	if c1.phy > c2.phy {
		fmt.Println("Meglio", c1.nome, "con", c1.phy, "di fisico")
	} else {
		fmt.Println("Meglio", c2.nome, "con", c2.phy, "di fisico")
	}
	fmt.Println("Totale:", c1.nome, (c1.vel + c1.tir + c1.pas + c1.dri + c1.def + c1.phy), "VS", c2.nome, c2.vel+c2.tir+c2.pas+c2.dri+c2.def+c2.phy)
}
