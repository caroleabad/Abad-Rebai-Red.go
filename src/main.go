package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	argent     int
}

func (p1 *personage) init(name string, class string, level int, viemax int, vieactuel int, sort string, argent int) {
	p1.name = name
	p1.class = class
	p1.level = level
	p1.viemax = viemax
	p1.vieactuel = vieactuel
	p1.skill = append(p1.skill, sort)
	p1.argent = argent
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
	//var p1 personage
	/*
		p.init("caROle", "Elfe", 1, 100, 80, map[string]int{"potions": 1, "clefs": 5, "hache": 1, "potions de poison": 1}, []string{"coup de poing"})
		p1.init("Rihem", "humain,elfe,nain", 1, 100, 40, map[string]int{"boule de feu": 3, " potion": 1, "épée": 1, "bouclier": 1}, []string{"voler dans les airs"})
		p.charCreation()
		p.Menu()*/
	p.initPerso()

	/*  println("1 Jouer")
	println("2 Marchand")
	println("3 Quitter")
	time.Sleep(1 * time.Second)
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
		} else if reponse == "3" {
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
	} */
}

func (p personage) Menu() {
	out := false
	for !out {
		fmt.Println("1 Afficher les informations du joueur")
		fmt.Println("2 Afficher inventaire")
		fmt.Println("3 Afficher le marchand")
		fmt.Println("4 Afficher monstre")
		fmt.Println("5 Quitter")

		time.Sleep(1 * time.Second)
		input := WaitForInput()

		switch input {
		case "1":
			p.DisplayInfo()
			p.Menu()
		case "2":
			p.AccessInventory()
			p.Menu()
		case "3":
			p.Marchand()
			p.Menu()
		case "4":
			p.monstre()
			p.Menu()
		case "5":
			fmt.Println("Quitter")
			out = true

		default:
			fmt.Println("Mauvaise réponse")
		}
	}
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
	for i, v := range listeInventaire {
		fmt.Println((i + 1), " ", v.name, " ", v.price, "$")
	}
	var nb int
	var isCorrect bool
	for !isCorrect {
		var err error
		input := WaitForInput()
		nb, err = strconv.Atoi(input)
		if err == nil && nb > 0 && nb <= len(listeInventaire) {
			isCorrect = true
		} else {
			fmt.Println("il faut rentrer un nombre")
		}
	}

	if p.argent >= listeInventaire[nb-1].price {
		fmt.Println("vous avez acheter quelque chose")
		p.argent -= listeInventaire[nb-1].price
		p.inventaire[listeInventaire[nb-1].name] += 1
	} else {
		fmt.Println("vous n'avez pas assez d'argent ")

	}

}

//fmt.Println("Que vouliez-vous acheter ?")
//fmt.Print("1 Potion")
//fmt.Println("2 clef")
//fmt.Println("3 hache")
//fmt.Println("4 boule de feu")
//reponse := WaitForInput()
//if reponse == "1" {
//	p.inventaire["potions"]++
//	fmt.Println("vous avez bien acheter")
//} else if reponse == "2" {
//	p.inventaire["clefs"]++
//	fmt.Println("vous avez bien acheter")
//} else if reponse == "3" {
//	p.inventaire["hache"]++
//	fmt.Println("vous avez bien acheter")
//} else if reponse == "4" {
//	p.inventaire["boule de feu"]++
//	fmt.Println("vous avez bien acheter")
//}

func (p1 *personage) charCreation() {

	a := ""
	for i, v := range p1.name {
		if i == 0 {
			if 97 <= v && v <= 122 {
				a += string(v - 32)
			} else {
				a += string(v)
			}
		} else {
			if 65 <= v && v <= 90 {
				a += string(v + 32)
			} else {
				a += string(v)
			}
		}
	}
	p1.name = a

}
func (p *personage) initPerso() {
	fmt.Println("entrez votre nom:")
	var name string
	fmt.Scan(&name)
	var pviemax int
	var pvieactuel int
	var class string
	var level int
	var sort string
	fmt.Println("choisissez votre classe:")
	fmt.Scan(&class)

	switch class {
	case "humain":
		class = "humain"
		pviemax = 100
		pvieactuel = 50
		level = 1
		sort = "coup de poing"

	case "elfe":
		class = "elfe"
		pviemax = 80
		pvieactuel = 40
		level = 1
		sort = "coup de poing"

	case "nain":
		class = "nain"
		pviemax = 120
		pvieactuel = 60
		level = 1
		sort = "coup de poing"

	default:
		println("classe invalide, choisissez parmi humain , elfe ou nain ")
	}
	p.init(name, class, level, pviemax, pvieactuel, sort, 100)
	p.charCreation()
	p.Menu()

}

func (p *personage) ajouterObjet(objet string) {

	if len(p.inventaire) <= 10 || p.inventaire[objet] > 0 {
		p.inventaire[objet] += 1
		fmt.Println("a été ajouté à l'inventaire.\n", objet)
	} else {
		fmt.Println("L'inventaire est plein, vous ne pouvez pas ajouter plus d'objets.")
	}
}

type product struct {
	name  string
	price int
}

var potionDeVie = product{
	name:  "Potion de vie",
	price: 3,
}
var potionDePoison = product{
	name:  "potion de poison ",
	price: 6,
}
var livreDeSort = product{
	name:  "livre de sort: boule de feu ",
	price: 25,
}
var fourrure = product{
	name:  "fourrure de loup",
	price: 4,
}
var PeauDeTroll = product{
	name:  "peau de troll",
	price: 7,
}
var Cuir = product{
	name:  "cuir de sanglier",
	price: 3,
}
var Plume = product{
	name:  "plume de corbeau ",
	price: 1,
}
var Chapeau = product{
	name:  "chapeau de l'aventurier ",
	price: 5,
}
var Tunique = product{
	name:  " tunique de l'aventurier ",
	price: 5,
}
var Bottes = product{
	name:  " bottes de l'aventurier",
	price: 5,
}
var listeInventaire = []product{
	potionDeVie,
	potionDePoison,
	livreDeSort,
	fourrure,
	PeauDeTroll,
	Cuir,
	Plume,
	Chapeau,
	Tunique,
	Bottes,
}

type monstre struct {
	name          string
	pvmax         int
	pvactuel      int
	pointsattaque int
}

func (p *monstre) Init(name string, pvmax int, pvactuel int, pointsattaque int) {
	p.name = name
	p.pvmax = pvmax
	p.pvactuel = pvactuel
	p.pointsattaque = pointsattaque
}
func (p *monstre) InitGoblin() {
	fmt.Println("mon nom sera :", p.name)
	fmt.Println("mes points de vie maximum sont:", p.pvmax)
	fmt.Println("mes points de vie actuels sont:", p.pvactuel)
	fmt.Println("mes points d'attaque sont :", p.pointsattaque)
}
