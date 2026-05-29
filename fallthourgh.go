package main

import "fmt"

type Forfait int

const (
	Base Forfait = iota
	Confort
	Premium
)

type Vehicule struct {
	Immat   string
	Forfait Forfait
}

func prestations(f Forfait) []string {
	services := []string{}

	switch f {
	case Premium:
		services = append(services, "Véhicule de remplacement")
		fallthrough
	case Confort:
		services = append(services, "Révision annuelle")
		fallthrough
	case Base:
		services = append(services, "Contrôle technique")
	}

	return services
}

func main() {
	flotte := []Vehicule{
		{"AB-123-CD", Base},
		{"EF-456-GH", Confort},
		{"IJ-789-KL", Premium},
	}

	noms := []string{"Base", "Confort", "Premium"}

	premium := []string{}

	for _, v := range flotte {
		fmt.Printf("\nVéhicule %s — forfait %s :\n", v.Immat, noms[v.Forfait])

		services := prestations(v.Forfait)

		for i, s := range services {
			fmt.Printf("  %d. %s\n", i+1, s)
		}

		if v.Forfait == Premium {
			premium = append(premium, v.Immat)
		}
	}

	fmt.Println("\nVéhicules Premium :", premium)
}
