package main

import (
	"errors"
	"fmt"
)

func operer(a float64, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division par zéro")
		}
		return a / b, nil
	default:
		return 0, errors.New("opération inconnue")
	}
}

func creerOperation(op string) func(float64, float64) float64 {
	return func(a, b float64) float64 {
		res, err := operer(a, b, op)
		if err != nil {
			fmt.Println("Erreur :", err)
			return 0
		}
		return res
	}
}

func main() {
	var a, b float64
	var op string

	for {
		fmt.Print("Entrez deux nombres et une opération (ou 'quit') : ")
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&op)

		if op == "quit" {
			break
		}

		res, err := operer(a, b, op)
		if err != nil {
			fmt.Println("Erreur :", err)
		} else {
			fmt.Println("Résultat :", res)
		}
	}
}
