package main

import (
	"fmt"
	"math"
)

const (
	IMCMaigreur = 18.5
	IMCNormal   = 25.0
	IMCSurpoids = 30.0
)

func main() {
	var poids float64
	var taille float64

	fmt.Print("Entrez votre poids : ")
	fmt.Scan(&poids)

	fmt.Print("Entrez votre taille : ")
	fmt.Scan(&taille)

	imc := math.Round(poids/(taille*taille)*100) / 100

	fmt.Println("Poids :", poids, "kg")
	fmt.Println("Taille :", taille, "m")
	fmt.Println("IMC :", imc)

	var categorie string
	switch {
	case imc < IMCMaigreur:
		categorie = "Maigreur"
	case imc < IMCNormal:
		categorie = "Normal"
	case imc < IMCSurpoids:
		categorie = "Surpoids"
	default:
		categorie = "Obésité"
	}

	fmt.Println("Catégorie :", categorie)
}
