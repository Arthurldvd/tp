package main

import "fmt"

type Personne struct {
	Prenom string
	Nom    string
	Age    int
	Email  string
}

func (p Personne) NomComplet() string {
	return p.Prenom + " " + p.Nom
}

func (p Personne) Presentation() string {
	return fmt.Sprintf("%s, %d ans (%s)", p.NomComplet(), p.Age, p.Email)
}

type Adresse struct {
	Rue        string
	Ville      string
	CodePostal string
}

func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

type Employe struct {
	Personne
	Adresse
	Poste   string
	Salaire float64
}

func (e Employe) FicheEmploye() string {
	return fmt.Sprintf(
		"Employé : %s\n  Poste   : %s\n  Salaire : %.2f €\n  Adresse : %s",
		e.Presentation(), e.Poste, e.Salaire, e.Format(),
	)
}

func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire += e.Salaire * pct / 100
}

type Etudiant struct {
	Personne
	Promo   string
	Moyenne float64
}

func (et Etudiant) MentionObtenue() string {
	switch {
	case et.Moyenne >= 16:
		return "Très Bien"
	case et.Moyenne >= 14:
		return "Bien"
	case et.Moyenne >= 12:
		return "Assez Bien"
	case et.Moyenne >= 10:
		return "Passable"
	default:
		return "Insuffisant"
	}
}

func (et Etudiant) FicheEtudiant() string {
	return fmt.Sprintf(
		"Étudiant : %s\n  Promo   : %s\n  Moyenne : %.2f (%s)",
		et.Presentation(), et.Promo, et.Moyenne, et.MentionObtenue(),
	)
}

func main() {
	employes := []Employe{
		{
			Personne: Personne{"Alice", "Martin", 34, "alice.martin@wanadev.fr"},
			Adresse:  Adresse{"12 rue des Lilas", "Lyon", "69003"},
			Poste:    "Développeuse",
			Salaire:  3200,
		},
		{
			Personne: Personne{"Bruno", "Lefevre", 41, "bruno.lefevre@wanadev.fr"},
			Adresse:  Adresse{"5 avenue Jean Jaurès", "Villeurbanne", "69100"},
			Poste:    "Chef de projet",
			Salaire:  3800,
		},
	}

	etudiants := []Etudiant{
		{
			Personne: Personne{"Chloé", "Dubois", 21, "maxence.dubois@esgi.fr"},
			Promo:    "M1 Informatique",
			Moyenne:  15.5,
		},
		{
			Personne: Personne{"David", "Nguyen", 22, "axel.rouqette@esgi.fr"},
			Promo:    "M2 Informatique",
			Moyenne:  11.2,
		},
	}

	employes[0].AugmenterSalaire(5)

	fmt.Println("------ Employés ------")
	for _, e := range employes {
		fmt.Println(e.FicheEmploye())
		fmt.Println()
	}

	fmt.Println("------ Étudiants ------")
	for _, et := range etudiants {
		fmt.Println(et.FicheEtudiant())
		fmt.Println()
	}
}
