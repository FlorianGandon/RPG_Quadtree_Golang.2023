package quadtree

import (
	"fmt"
	"testing"

	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/tiles"
)

func sont_different(t1 [][]tiles.Tiles, t2 [][]tiles.Tiles) bool {
	if len(t1) != len(t2) {
		return true
	}
	for i := 0; i < len(t1); i++ {
		if len(t1[i]) != len(t2[i]) {
			return true
		}
		for j := 0; j < len(t1[i]); j++ {
			if t1[i][j].Types != t2[i][j].Types {
				return true
			}
		}
	}
	return false
}

func Test_UpdateSquare(t *testing.T) {
	var fullContent_init [][]tiles.Tiles = [][]tiles.Tiles{
		{{Types: 0}, {Types: 0}, {Types: 1}, {Types: 0}, {Types: 2}, {Types: 2}, {Types: 2}, {Types: 2}},
		{{Types: 0}, {Types: 0}, {Types: 1}, {Types: 0}, {Types: 2}, {Types: 2}, {Types: 2}, {Types: 2}},
		{{Types: 0}, {Types: 0}, {Types: 1}, {Types: 0}, {Types: 2}, {Types: 2}, {Types: 2}, {Types: 2}},
		{{Types: 0}, {Types: 0}, {Types: 1}, {Types: 1}, {Types: 0}, {Types: 0}, {Types: 0}, {Types: 0}},
		{{Types: 0}, {Types: 0}, {Types: 0}, {Types: 1}, {Types: 1}, {Types: 1}, {Types: 0}, {Types: 0}},
		{{Types: 2}, {Types: 2}, {Types: 0}, {Types: 0}, {Types: 0}, {Types: 1}, {Types: 0}, {Types: 0}},
		{{Types: 2}, {Types: 2}, {Types: 2}, {Types: 0}, {Types: 0}, {Types: 1}, {Types: 0}, {Types: 0}},
		{{Types: 2}, {Types: 2}, {Types: 2}, {Types: 0}, {Types: 0}, {Types: 1}, {Types: 0}, {Types: 0}},
	}

	var fullContent [][]tiles.Tiles = [][]tiles.Tiles{
		{{Types: 0}, {Types: 0}, {Types: 1}, {Types: 0}, {Types: 2}, {Types: 2}, {Types: 2}, {Types: 2}},
		{{Types: 0}, {Types: 0}, {Types: 1}, {Types: 0}, {Types: 2}, {Types: 2}, {Types: 2}, {Types: 2}},
		{{Types: 0}, {Types: 0}, {Types: 1}, {Types: 0}, {Types: 2}, {Types: 2}, {Types: 2}, {Types: 2}},
		{{Types: 0}, {Types: 0}, {Types: 1}, {Types: 1}, {Types: 0}, {Types: 0}, {Types: 0}, {Types: 0}},
		{{Types: 0}, {Types: 0}, {Types: 0}, {Types: 1}, {Types: 1}, {Types: 1}, {Types: 0}, {Types: 0}},
		{{Types: 2}, {Types: 2}, {Types: 0}, {Types: 0}, {Types: 0}, {Types: 1}, {Types: 0}, {Types: 0}},
		{{Types: 2}, {Types: 2}, {Types: 2}, {Types: 0}, {Types: 0}, {Types: 1}, {Types: 0}, {Types: 0}},
		{{Types: 2}, {Types: 2}, {Types: 2}, {Types: 0}, {Types: 0}, {Types: 1}, {Types: 0}, {Types: 0}},
	}
	var q Quadtree = MakeFromArray(fullContent)
	var content [][]tiles.Tiles
	var content_voulu [][]tiles.Tiles

	if sont_different(fullContent_init, fullContent) {
		fmt.Println("modification du full content")
		t.Fail()
	}

	content_voulu = fullContent_init
	// content ne se "vide" avant l'appel pour une fonctionnalité
	content = [][]tiles.Tiles{
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
	}
	q.GetContent(1, 1, content) //1,1 correspond à 0,0 ... donc je veux afficher ici tout le tableau
	if sont_different(content_voulu, content) {
		fmt.Println("echec de la récupération de la carte en entier")
		t.Fail()
	}

	content_voulu = [][]tiles.Tiles{
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
	}
	// content ne se "vide" avant l'appel pour une fonctionnalité
	content = [][]tiles.Tiles{
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
	}
	q.GetContent(-9, -9, content) //-9,-9 correspond à -10,-10 ... donc je veux afficher ici rien
	if sont_different(content_voulu, content) {
		fmt.Println("echec de la récupération du content en dehors de la carte")
		t.Fail()
	}

	content_voulu = [][]tiles.Tiles{
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: 0}, {Types: 0}, {Types: 1}, {Types: 0}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: 0}, {Types: 0}, {Types: 1}, {Types: 0}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: 0}, {Types: 0}, {Types: 1}, {Types: 0}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: 0}, {Types: 0}, {Types: 1}, {Types: 1}},
	}
	// content ne se "vide" avant l'appel pour une fonctionnalité
	content = [][]tiles.Tiles{
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
		{{Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}, {Types: -1}},
	}
	q.GetContent(-3, -3, content) //-3,-3 correspond à -4,-4 ... donc je veux afficher 1/4 de la carte
	if sont_different(content_voulu, content) {
		fmt.Println("echec de la récupération du content à 1/4 dans la carte")
		t.Fail()
	}

}

func Test_UpdateRectangle(t *testing.T) {
	var fullContent_init [][]tiles.Tiles = [][]tiles.Tiles{
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

	var fullContent [][]tiles.Tiles = [][]tiles.Tiles{
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
	var q Quadtree = MakeFromArray(fullContent)
	var content [][]tiles.Tiles
	var content_voulu [][]tiles.Tiles

	if sont_different(fullContent_init, fullContent) {
		fmt.Println("modification du full content")
		t.Fail()
	}

	content_voulu = fullContent_init
	// content ne se "vide" avant l'appel pour une fonctionnalité
	content = [][]tiles.Tiles{
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
	q.GetContent(1, 1, content) //1,1 correspond à 0,0 ... donc je veux afficher ici tout le tableau
	if sont_different(content_voulu, content) {
		fmt.Println("echec de la récupération de la carte en entier")
		t.Fail()
	}

	content_voulu = [][]tiles.Tiles{
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
	content = [][]tiles.Tiles{
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
	q.GetContent(-9, -9, content) //-9,-9 correspond à -10,-10 ... donc je veux afficher ici rien
	if sont_different(content_voulu, content) {
		fmt.Println("echec de la récupération du content en dehors de la carte")
		t.Fail()
	}

	content_voulu = [][]tiles.Tiles{
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
	// content ne se "vide" avant l'appel pour une fonctionnalité
	content = [][]tiles.Tiles{
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
	q.GetContent(-2, -3, content) //-2,-3 correspond à -3,-4 ... donc je veux afficher environ 1/4 de la carte
	if sont_different(content_voulu, content) {
		fmt.Println("echec de la récupération du content à environ 1/4 dans la carte")
		t.Fail()
	}

}

func Test_estCarreIdentique(t *testing.T) {
	var fullContent [][]tiles.Tiles = [][]tiles.Tiles{
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

	if estCarreIdentique(element{n: &node{topLeftX: 0, topLeftY: 0, width: 2, height: 2}}, fullContent) {
		t.Fail()
	}

	if !estCarreIdentique(element{n: &node{topLeftX: 0, topLeftY: 0, width: 1, height: 2}}, fullContent) {
		t.Fail()
	}

	if estCarreIdentique(element{n: &node{topLeftX: 0, topLeftY: 0, width: 0, height: 2}}, fullContent) {
		t.Fail()
	}

	if estCarreIdentique(element{n: &node{topLeftX: 0, topLeftY: 0, width: 2, height: 0}}, fullContent) {
		t.Fail()
	}

	if !estCarreIdentique(element{n: &node{topLeftX: 0, topLeftY: 5, width: 2, height: 4}}, fullContent) {
		t.Fail()
	}
}
