package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)
//tests

type personage struct {
	name       string
	class      string
	level      int
	viemax     int
	vieactuel  int
	inventaire map[string]int
}

func main() {

	var p personage

	p.name = "Carole"
	p.class = "Elfe"
	p.level = 1
	p.viemax = 100
	p.vieactuel = 40
	p.inventaire = map[string]int{"potions": 3, "clefs": 5, "hache": 1}

	println("1 Jouer")
	println("2 Marchand")
	println("3 Quitter")
	input := WaitForInput()

	switch input {
	case "1":
		println("Jouer")
		break
	case "2":
		println("Que voulez-vous acheter ?")
		println("1 Potion")
		println("2 clef")
		println("3 hache")
		reponse := WaitForInput()
		if reponse == "1" {
			p.inventaire["potions"]++
			println("vous avez bien acheter")
		} else if reponse == "2" {
			p.inventaire["clefs"]++
			println("vous avez bien acheter")
		}else if reponse == "3"{
			p.inventaire["hache"]++
			println("vous avez bien acheter")
		}
		
		break
	case "3":
		println("Quitter")
		break
	default:
		println("Mauvaise réponse")
		break
	}
}

func (p personage) DisplayInfo() {

	fmt.Println("---------------")
	fmt.Println("je m'appele", p.name)
	fmt.Println("ma classe est", p.class)
	fmt.Println("mon niveau est", p.level)
	fmt.Println("j'ai une vie de maximum", p.viemax)
	fmt.Println("J'en est actuelement", p.vieactuel)
	fmt.Println("j'ai dans l'inventaire", p.inventaire)
}

func (p personage) AccessInventory() {

	fmt.Println("Dans mon inventaire j'ai")
	for clef, valeur := range p.inventaire {
		fmt.Println(valeur, "   ", clef)
	}
}

func (p *personage) TakePot() {
	if p.inventaire["potions"] > 0 {
		p.inventaire["potions"]--

		p.vieactuel = 90
		fmt.Println("Vie actuelle : ", p.vieactuel)
	}
}

func WaitForInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}
