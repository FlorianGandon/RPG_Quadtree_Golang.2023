package floor

import (
	"fmt"
	"testing"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/tiles"
)

func sont_identique2(t1 [][]tiles.Tiles, t2 [][]tiles.Tiles) bool {
	if len(t1) != len(t2) {
		return false
	}
	for i := 0; i < len(t1); i++ {
		if len(t1[i]) != len(t2[i]) {
			return false
		}
		for j := 0; j < len(t1[i]); j++ {
			if t1[i][j].Types != t2[i][j].Types {
				return false
			}
		}
	}
	return true
}

func Test_Update(t *testing.T) {
	var f_beaupasbeau Floor
	f_beaupasbeau.fullContent = readFloorFromFile("../floor-files/beaupasbeau")
	f_beaupasbeau.Height = len(f_beaupasbeau.fullContent)
	f_beaupasbeau.Width = len(f_beaupasbeau.fullContent[0])
	f_beaupasbeau.Content = make([][]tiles.Tiles, 8)
	for y := 0; y < len(f_beaupasbeau.Content); y++ {
		f_beaupasbeau.Content[y] = make([]tiles.Tiles, 8)
	}

	var tab_voulu [][]tiles.Tiles = readFloorFromFile("../floor-files/beaupasbeau")
	f_beaupasbeau.updateFromFileFloor(1, 1) //1,1 correspond à 0,0 ... donc je veux afficher ici tout le tableau
	if !sont_identique2(f_beaupasbeau.Content, tab_voulu) {
		fmt.Println("echec de la récupération de la carte en entier")
		t.Fail()
	}

	tab_voulu = [][]tiles.Tiles{{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
	}
	f_beaupasbeau.updateFromFileFloor(-9, -9) //-9,-9 correspond à -10,-10 ... donc je veux afficher ici rien
	if !sont_identique2(f_beaupasbeau.Content, tab_voulu) {
		fmt.Println("echec de la récupération du content en dehors de la carte")
		t.Fail()
	}

	tab_voulu = [][]tiles.Tiles{
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: 0}, {Types: 0}, {Types: 1}, {Types: 0}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: 0}, {Types: 0}, {Types: 1}, {Types: 0}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: 0}, {Types: 0}, {Types: 1}, {Types: 0}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: 0}, {Types: 0}, {Types: 1}, {Types: 1}},
	}
	f_beaupasbeau.updateFromFileFloor(-3, -3) //-3,-3 correspond à -4,-4 ... donc je veux afficher ici rien
	if !sont_identique2(f_beaupasbeau.Content, tab_voulu) {
		fmt.Println("echec de la récupération du content à 1/4 dans la carte")
		t.Fail()
	}
}

func Test_UpdateRectangle(t *testing.T) {
	var f_beaupasbeau Floor
	var tab_voulu [][]tiles.Tiles

	f_beaupasbeau.fullContent = [][]tiles.Tiles{
		{{Types: 1}, {Types: 0}, {Types: 2}, {Types: 2}, {Types: 2}, {Types: 3}, {Types: 2}},
		{{Types: 1}, {Types: 0}, {Types: 2}, {Types: 2}, {Types: 3}, {Types: 2}, {Types: 2}},
		{{Types: 0}, {Types: 1}, {Types: 0}, {Types: 2}, {Types: 3}, {Types: 2}, {Types: 2}},
		{{Types: 0}, {Types: 1}, {Types: 2}, {Types: 1}, {Types: 0}, {Types: 4}, {Types: 0}},
		{{Types: 0}, {Types: 1}, {Types: 1}, {Types: 1}, {Types: 1}, {Types: 0}, {Types: 8}},
		{{Types: 2}, {Types: 2}, {Types: 0}, {Types: 0}, {Types: 6}, {Types: 0}, {Types: 8}},
		{{Types: 2}, {Types: 2}, {Types: 0}, {Types: 3}, {Types: 6}, {Types: 4}, {Types: 0}},
		{{Types: 2}, {Types: 2}, {Types: 0}, {Types: 3}, {Types: 6}, {Types: 0}, {Types: 8}},
		{{Types: 2}, {Types: 2}, {Types: 0}, {Types: 3}, {Types: 6}, {Types: 4}, {Types: 0}},
	}
	f_beaupasbeau.Height = len(f_beaupasbeau.fullContent)
	f_beaupasbeau.Width = len(f_beaupasbeau.fullContent[0])
	f_beaupasbeau.Content = make([][]tiles.Tiles, 9)
	for y := 0; y < len(f_beaupasbeau.Content); y++ {
		f_beaupasbeau.Content[y] = make([]tiles.Tiles, 7)
	}

	tab_voulu = f_beaupasbeau.fullContent
	f_beaupasbeau.updateFromFileFloor(1, 1) //1,1 correspond à 0,0 ... donc je veux afficher ici tout le tableau
	if !sont_identique2(f_beaupasbeau.Content, tab_voulu) {
		fmt.Println("echec de la récupération de la carte en entier")
		t.Fail()
	}

	tab_voulu = [][]tiles.Tiles{
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
	}
	// content ne se "vide" avant l'appel pour une fonctionnalité
	f_beaupasbeau.updateFromFileFloor(-9, -9) //-9,-9 correspond à -10,-10 ... donc je veux afficher ici rien
	if !sont_identique2(f_beaupasbeau.Content, tab_voulu) {
		fmt.Println("echec de la récupération du content en dehors de la carte")
		t.Fail()
	}

	tab_voulu = [][]tiles.Tiles{
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: 1}, {Types: 0}, {Types: 2}, {Types: 2}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: 1}, {Types: 0}, {Types: 2}, {Types: 2}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: 0}, {Types: 1}, {Types: 0}, {Types: 2}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: 0}, {Types: 1}, {Types: 2}, {Types: 1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: 0}, {Types: 1}, {Types: 1}, {Types: 1}},
	}
	f_beaupasbeau.updateFromFileFloor(-2, -3) //-2,-3 correspond à -3,-4 ... donc je veux afficher environ 1/4 de la carte
	if !sont_identique2(f_beaupasbeau.Content, tab_voulu) {
		fmt.Println("echec de la récupération du content à environ 1/4 dans la carte")
		t.Fail()
	}
}
