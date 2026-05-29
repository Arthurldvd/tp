package main

import (
	"errors"
	"fmt"
	"strings"
)

type Produit struct {
	ID        int
	Nom       string
	Marque    string
	Prix      float64
	Stock     int
	Categorie string
	Actif     bool
}

type Catalogue struct {
	produits   []Produit
	prochainID int
}

func (c *Catalogue) AjouterProduit(p Produit) int {
	if c.prochainID == 0 {
		c.prochainID = 1
		for _, existant := range c.produits {
			if existant.ID >= c.prochainID {
				c.prochainID = existant.ID + 1
			}
		}
	}
	p.ID = c.prochainID
	c.prochainID++
	c.produits = append(c.produits, p)
	return p.ID
}

func (c Catalogue) TrouverParID(id int) (Produit, error) {
	for _, p := range c.produits {
		if p.ID == id {
			return p, nil
		}
	}
	return Produit{}, errors.New("produit introuvable")
}

func (c Catalogue) TrouverParCategorie(cat string) []Produit {
	resultats := []Produit{}
	for _, p := range c.produits {
		if strings.EqualFold(p.Categorie, cat) {
			resultats = append(resultats, p)
		}
	}
	return resultats
}

func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	nb := 0
	for i := range c.produits {
		if strings.EqualFold(c.produits[i].Categorie, categorie) {
			c.produits[i].Prix -= c.produits[i].Prix * pct / 100
			nb++
		}
	}
	return nb
}

func (c *Catalogue) Vendre(id int, qte int) error {
	for i := range c.produits {
		if c.produits[i].ID == id {
			if c.produits[i].Stock < qte {
				return fmt.Errorf("stock insuffisant (%d disponible(s))", c.produits[i].Stock)
			}
			c.produits[i].Stock -= qte
			return nil
		}
	}
	return errors.New("produit introuvable")
}

func (c Catalogue) Rapport() string {
	var valeurTotale float64
	for _, p := range c.produits {
		valeurTotale += p.Prix * float64(p.Stock)
	}
	return fmt.Sprintf("Catalogue : %d produit(s) — valeur totale du stock : %.2f €",
		len(c.produits), valeurTotale)
}

func main() {
	catalogue := Catalogue{
		produits: []Produit{
			{1, "iPhone 15 Pro", "Apple", 1229.00, 12, "Smartphone", true},
			{2, "MacBook Air M3", "Apple", 1499.00, 7, "Ordinateur", true},
			{3, "Galaxy S24", "Samsung", 899.00, 15, "Smartphone", true},
			{4, "ThinkPad X1", "Lenovo", 1799.00, 4, "Ordinateur", true},
			{5, "Casque WH-1000XM5", "Sony", 379.00, 25, "Audio", true},
		},
	}

	for {
		fmt.Println("\n===== TechShop =====")
		fmt.Println("[1] Ajouter  [2] Chercher  [3] Soldes  [4] Vendre  [5] Rapport  [0] Quitter")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			var p Produit
			fmt.Print("Nom : ")
			fmt.Scan(&p.Nom)
			fmt.Print("Marque : ")
			fmt.Scan(&p.Marque)
			fmt.Print("Prix : ")
			fmt.Scan(&p.Prix)
			fmt.Print("Stock : ")
			fmt.Scan(&p.Stock)
			fmt.Print("Catégorie : ")
			fmt.Scan(&p.Categorie)
			p.Actif = true

			id := catalogue.AjouterProduit(p)
			fmt.Printf("Produit ajouté (ID %d).\n", id)

		case 2:
			fmt.Print("ID à chercher : ")
			var id int
			fmt.Scan(&id)

			p, err := catalogue.TrouverParID(id)
			if err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Printf("%s (%s) — %.2f € — stock : %d — %s\n",
					p.Nom, p.Marque, p.Prix, p.Stock, p.Categorie)
			}

		case 3:
			fmt.Print("Catégorie en solde : ")
			var cat string
			fmt.Scan(&cat)
			fmt.Print("Pourcentage de réduction : ")
			var pct float64
			fmt.Scan(&pct)

			nb := catalogue.AppliquerReduction(cat, pct)
			fmt.Printf("%d produit(s) en promotion.\n", nb)

		case 4:
			fmt.Print("ID du produit vendu : ")
			var id int
			fmt.Scan(&id)
			fmt.Print("Quantité : ")
			var qte int
			fmt.Scan(&qte)

			if err := catalogue.Vendre(id, qte); err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Vente enregistrée.")
			}

		case 5:
			fmt.Println(catalogue.Rapport())

		case 0:
			fmt.Println("À bientôt chez TechShop !")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}
