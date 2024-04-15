package tiles

// Tiles est utilisée dans le but de la génération
type Tiles struct {
	Types           int           `json:"t"`
	Entropy         int           `json:"E"`
	ListPossibility []Possibility `json:"TE"`
}

type Possibility struct {
	Possibility int `json:"P"`
	Weight      int `json:"W"`
}

// ChangeWeight change le poids de ListPossibility en fonction d'une valeur et d'un poids a rajouté
func ChangeWeight(possibility []Possibility, value, weight int) {
	for i, y := range possibility {
		if y.Possibility == value {
			possibility[i].Weight += weight
			break
		}
	}
}

var _ = Tiles{4, -1, []Possibility{}} //Règle problème IDE
