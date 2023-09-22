package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type personage struct {
	name       string
	class      string
	level      int
	viemax     int
	vieactuel  int
	inventaire map[string]int
	skill      []string
}

func (p *personage) init(name string, class string, level int, viemax int, vieactuel int, inventaire map[string]int, skill []string) {
	p.name = name
	p.class = class
	p.level = level
	p.viemax = viemax
	p.vieactuel = vieactuel
	p.inventaire = inventaire
	p.skill = skill
} 

func (p *personage) potion() {
	if p.vieactuel+50 >= p.viemax {
		p.vieactuel = p.viemax
	} else {
		p.vieactuel += 50
	}

	fmt.Println("Vos PV sont de  : ", p.vieactuel, "/", p.viemax)
}

func main() {

	var p personage
	p.init("Carole", "Elfe", 1, 100, 80, map[string]int{"potions": 1, "clefs": 5, "hache": 1, "potions de poison": 1 , "boule de feu":1}, []string{"coup de poing"})
	p.DisplayInfo()
	p.PoisonPot()
	p.potion()
	p.spellBook()

	// println("1 Jouer")
	// println("2 Marchand")
	// println("3 Quitter")
	// time.Sleep(1 * time.Second)
	// input := WaitForInput()

	// switch input {
	// case "1":
	// 	println("Jouer")
	// 	break
	// case "2":
	// 	println("Que voulez-vous acheter ?")
	// 	println("1 Potion")
	// 	println("2 clef")
	// 	println("3 hache")
	// 	reponse := WaitForInput()
	// 	if reponse == "1" {
	// 		p.inventaire["potions"]++
	// 		println("vous avez bien acheter")
	// 	} else if reponse == "2" {
	// 		p.inventaire["clefs"]++
	// 		println("vous avez bien acheter")
	// 	} else if reponse == "3" {
	// 		p.inventaire["hache"]++
	// 		println("vous avez bien acheter")
	// 	}

	// 	break
	// case "3":
	// 	println("Quitter")
	// 	break
	// default:
	// 	println("Mauvaise réponse")
	// 	break
	// }
}

func (p *personage) DisplayInfo() {

	fmt.Println("---------------")
	fmt.Println("je m'appelle", p.name)
	fmt.Println("ma classe est", p.class)
	fmt.Println("mon niveau est", p.level)
	fmt.Println("j'ai une vie de maximum", p.viemax)
	fmt.Println("J'en ai actuellement", p.vieactuel)
	fmt.Println("j'ai dans l'inventaire", p.inventaire)
	fmt.Println("Compétences : ", p.skill)
}

func (p *personage) AccessInventory() {

	fmt.Println("Dans mon inventaire j'ai")
	for clef, valeur := range p.inventaire {
		fmt.Println(valeur, "   ", clef)
	}
}

func (p *personage) TakePot() {
	if p.inventaire["potions"] > 0 {
		fmt.Println("Vous prenez une potion de vie")
		p.inventaire["potions"]--
		if p.vieactuel+50 > p.viemax {
			p.vieactuel = p.viemax
		} else {
			p.vieactuel += 50
		}
		fmt.Println("PV : ", p.vieactuel, "/", p.viemax)
	} else {
		fmt.Println("Vous n'avez pas de potion de vie")
	}
}

func (p *personage) PoisonPot() {
	if p.inventaire["potions de poison"] > 0 {
		fmt.Println("Vous prenez une potion de poison")
		p.inventaire["potions de poison"]--
		for i := 0; i < 3; i++ {
			p.vieactuel -= 10
			fmt.Println("PV : ", p.vieactuel, "/", p.viemax)
			time.Sleep(1 * time.Second)
		}
	} else {
		fmt.Println("Vous n'avez pas de potion de poison")
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
func (p *personage) dead() {
	if p.vieactuel <= 0 {
		fmt.Println("Le joueur est mort !")
		p.vieactuel = p.viemax / 2
		fmt.Printf("Le joueur a été ressuscité avec 50% de ses points de vie maximum : points de vie.\n", p.vieactuel)
	}
}

func (p *personage) poisonPot() {

	dureePoison := 3
	degatsParSeconde := 10

	fmt.Printf("Le poison inflige %d points de dégâts par seconde pendant %d secondes.\n", degatsParSeconde, dureePoison)

	for i := 1; i <= dureePoison; i++ {
		p.vieactuel -= degatsParSeconde
		fmt.Printf("Seconde %d : Points de vie actuels : %d / Points de vie maximum : %d\n", i, p.vieactuel, p.viemax)
	}

	fmt.Println("Le poison a cessé de faire effet.")
}
func (p *personage) spellBook() {
	if p.CheckSpell("Boule de feu") {
		fmt.Println("sort deja present")
	} else {
		p.skill = append(p.skill, "Boule de feu")
		fmt.Println("sort boule de feu ajouter a la liste", p.skill)
	}

}

func (p *personage) CheckSpell(spellName string) bool {
	for _, valeur := range p.skill {
		if valeur == spellName {
			return true
		}
	}
	return false
}
func (p *personage) Marchand() {
	println("Que voulez-vous acheter ?")
	println("1 Potion")
	println("2 clef")
	println("3 hache")
	println( "4 boule de feu")
	reponse := WaitForInput()
	if reponse == "1" {
		p.inventaire["potions"]++
		println("vous avez bien acheter")
	} else if reponse == "2" {
		p.inventaire["clefs"]++
		println("vous avez bien acheter")
	} else if reponse == "3" {
		p.inventaire["hache"]++
		println("vous avez bien acheter")
	}else if reponse == "4"{
		p.inventaire["boule de feu"]++
		println("vous avez bien acheter")
	}
	
}
