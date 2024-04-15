package generation

import (
	"github.com/mroth/weightedrand"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/Coords"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/tiles"
)

// DontConnectwith Contraintre enmpechant un case à se générer à coter d'une autre case
var DontConnectwith = [][]int{{}, {5, 2, 7}, {1}, {}, {}, {1, 7}, {}, {5, 1}}
var Weigh = 20

// Generation Génère un tronçon de carte de size*size inspirée de l'agorithme "Wave fonction colapse"
func Generation(size int, h, b, g, d []tiles.Tiles) (generatedTiles [][]tiles.Tiles) {
	var chunk [][]tiles.Tiles
	for i := 0; i < size; i++ {
		chunk = append(chunk, []tiles.Tiles{})
		for y := 0; y < size; y++ {
			var tilesEntropy = []tiles.Possibility{{0, 2}, {1, 1}, {2, 0}, {3, 0}, {4, 0}, {5, 1}, {6, 1}, {7, 1}}
			chunk[i] = append(chunk[i], tiles.Tiles{Types: -1, Entropy: len(tilesEntropy), ListPossibility: tilesEntropy})
		}
	}
	var minCoord = FindMin(chunk)
	CalculateListPossibilityNeibourg(chunk, h, b, g, d)
	generatedTiles = generateTiles(chunk, minCoord)
	return
}

// generateTiles génère toute les tiles de manière recursive
func generateTiles(chunk [][]tiles.Tiles, mincoord Coords.Coords) [][]tiles.Tiles {
	//fmt.Println(mincoord)
	if mincoord.X == -1 {
		return chunk
	}
	chunk[mincoord.Y][mincoord.X] = tiles.Tiles{Types: Select(chunk[mincoord.Y][mincoord.X].ListPossibility)}
	CalculateListPossibility(chunk, mincoord)
	var newMinCoord = FindMin(chunk)
	return generateTiles(chunk, newMinCoord)
}

// Select Choisie une des possibilités de la case en focntion des possibilitées et de leurs poids
func Select(possibility []tiles.Possibility) int {
	var c = make([]weightedrand.Choice, 0, 7)
	for _, i := range possibility {
		c = append(c, weightedrand.Choice{Weight: uint(i.Weight), Item: i.Possibility})
	}
	var result, err = weightedrand.NewChooser(c...)
	// ... trouvé par moi-même grâce à la doc
	if err != nil {
		panic("Error with random")
	}
	//fmt.Println(possibility, result)
	return result.PickSource(configuration.Global.SeedRand).(int)
}

// CalculateListPossibilityNeibourg Calcule la liste des possibilitées des casses adjacente au chunks voisin déjà généré
func CalculateListPossibilityNeibourg(chunk [][]tiles.Tiles, h, b, g, d []tiles.Tiles) {
	for x, y := range h {
		for _, i := range DontConnectwith[y.Types] {
			if !isConnectable(chunk[0][x], y.Types) {
				chunk[0][x].ListPossibility = remove(chunk[0][x].ListPossibility, i)
				chunk[0][x].Entropy = CalculateEntropy(chunk[0][x])
				tiles.ChangeWeight(chunk[0][x].ListPossibility, y.Types, Weigh)
			}
		}
	}
	for x, y := range b {
		for _, i := range DontConnectwith[y.Types] {
			if !isConnectable(chunk[len(chunk)-1][x], y.Types) {
				chunk[len(chunk)-1][x].ListPossibility = remove(chunk[len(chunk)-1][x].ListPossibility, i)
				chunk[len(chunk)-1][x].Entropy = CalculateEntropy(chunk[len(chunk)-1][x])
				tiles.ChangeWeight(chunk[len(chunk)-1][x].ListPossibility, y.Types, Weigh)
			}
		}
	}
	for x, y := range g {
		for _, i := range DontConnectwith[y.Types] {
			if !isConnectable(chunk[x][0], y.Types) {
				chunk[x][0].ListPossibility = remove(chunk[x][0].ListPossibility, i)
				chunk[x][0].Entropy = CalculateEntropy(chunk[x][0])
				tiles.ChangeWeight(chunk[x][0].ListPossibility, y.Types, Weigh)
			}
		}
	}
	for x, y := range d {
		for _, i := range DontConnectwith[y.Types] {
			if !isConnectable(chunk[x][len(chunk)-1], y.Types) {
				chunk[x][len(chunk)-1].ListPossibility = remove(chunk[x][len(chunk)-1].ListPossibility, i)
				chunk[x][len(chunk)-1].Entropy = CalculateEntropy(chunk[x][len(chunk)-1])
				tiles.ChangeWeight(chunk[x][len(chunk)-1].ListPossibility, y.Types, Weigh)
			}
		}
	}
}

// CalculateListPossibility Recalcule la liste des possibilitées des casses adjacente aux coordonnées doné
func CalculateListPossibility(chunk [][]tiles.Tiles, coordChange Coords.Coords) {
	var val = chunk[coordChange.Y][coordChange.X].Types
	if len(DontConnectwith[val]) != 0 {
		for i := range DontConnectwith[val] {
			if coordChange.Y != 0 {
				if !isConnectable(chunk[coordChange.Y-1][coordChange.X], val) {
					chunk[coordChange.Y-1][coordChange.X].ListPossibility = remove(chunk[coordChange.Y-1][coordChange.X].ListPossibility, DontConnectwith[val][i])
					chunk[coordChange.Y-1][coordChange.X].Entropy = CalculateEntropy(chunk[coordChange.Y-1][coordChange.X])
				}
			}
			if coordChange.Y != len(chunk)-1 {
				if !isConnectable(chunk[coordChange.Y+1][coordChange.X], val) {
					chunk[coordChange.Y+1][coordChange.X].ListPossibility = remove(chunk[coordChange.Y+1][coordChange.X].ListPossibility, DontConnectwith[val][i])
					chunk[coordChange.Y+1][coordChange.X].Entropy = CalculateEntropy(chunk[coordChange.Y+1][coordChange.X])
				}
			}

			if coordChange.X-1 >= 0 {
				if !isConnectable(chunk[coordChange.Y][coordChange.X-1], val) {
					chunk[coordChange.Y][coordChange.X-1].ListPossibility = remove(chunk[coordChange.Y][coordChange.X-1].ListPossibility, DontConnectwith[val][i])
					chunk[coordChange.Y][coordChange.X-1].Entropy = CalculateEntropy(chunk[coordChange.Y][coordChange.X-1])
				}
			}
			if coordChange.X+1 < len(chunk) {
				if !isConnectable(chunk[coordChange.Y][coordChange.X+1], val) {
					chunk[coordChange.Y][coordChange.X+1].ListPossibility = remove(chunk[coordChange.Y][coordChange.X+1].ListPossibility, DontConnectwith[val][i])
					chunk[coordChange.Y][coordChange.X+1].Entropy = CalculateEntropy(chunk[coordChange.Y][coordChange.X+1])
				}
			}
		}
		if val != 2 {
			if coordChange.Y != 0 && chunk[coordChange.Y-1][coordChange.X].Entropy != 0 {
				tiles.ChangeWeight(chunk[coordChange.Y-1][coordChange.X].ListPossibility, val, Weigh)
			} else if coordChange.Y != len(chunk)-1 && chunk[coordChange.Y+1][coordChange.X].Entropy != 0 {
				tiles.ChangeWeight(chunk[coordChange.Y+1][coordChange.X].ListPossibility, val, Weigh)
			} else if coordChange.X > 0 && chunk[coordChange.Y][coordChange.X-1].Entropy != 0 {
				tiles.ChangeWeight(chunk[coordChange.Y][coordChange.X-1].ListPossibility, val, Weigh)
			} else if coordChange.X+1 < len(chunk) && chunk[coordChange.Y][coordChange.X+1].Entropy != 0 {
				tiles.ChangeWeight(chunk[coordChange.Y][coordChange.X+1].ListPossibility, val, Weigh)
				//chunk[coordChange.Y][coordChange.X+1].Entropy = CalculateEntropy(chunk[coordChange.Y][coordChange.X+1])
			}
		}
	}
}

// isConnectable regarde si deux case peuvent se connecter entre elle
func isConnectable(tile tiles.Tiles, tile2 int) bool {
	for i := range tile.ListPossibility {
		if len(DontConnectwith[tile.ListPossibility[i].Possibility]) != 0 {
			for y := range DontConnectwith[tile.ListPossibility[i].Possibility] {
				if tile2 == DontConnectwith[tile.ListPossibility[i].Possibility][y] {
					return false
				}
			}

		}
	}
	return true
}

// FindMin trouve les entropies minimal et en choisie une aléatoirement
func FindMin(chunk [][]tiles.Tiles) (minCoords Coords.Coords) {
	var minimum = 100 //très moche
	var size = len(chunk)
	var listMinCoords []Coords.Coords
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if chunk[y][x].Entropy != 0 && minimum > chunk[y][x].Entropy {
				minimum = chunk[y][x].Entropy
				listMinCoords = []Coords.Coords{{x, y}}
			} else if chunk[y][x].Entropy == minimum {
				listMinCoords = append(listMinCoords, Coords.Coords{x, y})
			}
		}
	}
	if minimum == 100 {
		return Coords.Coords{X: -1, Y: -1}
	}
	return listMinCoords[configuration.Global.SeedRand.Intn(len(listMinCoords))]
}

// CalculateEntropy renvoie l'entropie d'une case
func CalculateEntropy(tilesEntropy tiles.Tiles) int {
	return len(tilesEntropy.ListPossibility)
}

// remove enlève un element dans une possibilité
func remove(Possibility []tiles.Possibility, elementtoremove int) []tiles.Possibility {
	for i, value := range Possibility {
		if value.Possibility == elementtoremove {
			return append(Possibility[:i], Possibility[i+1:]...)
		}
	}
	return Possibility
}
