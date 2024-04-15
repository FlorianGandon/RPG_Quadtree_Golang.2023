package floor

import (
	"fmt"
	"testing"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/tiles"
)

func sont_identique(t1 [][]int, t2 [][]tiles.Tiles) bool {
	if len(t1) != len(t2) {
		return false
	}
	for i := 0; i < len(t1); i++ {
		if len(t1[i]) != len(t2[i]) {
			return false
		}
		for j := 0; j < len(t1[i]); j++ {
			if t1[i][j] != t2[i][j].Types {
				return false
			}
		}
	}
	return true
}

func Test_beaupasbeau(t *testing.T) {
	var result [][]tiles.Tiles
	var result_voulu [][]int = [][]int{{0, 0, 1, 0, 2, 2, 2, 2}, {0, 0, 1, 0, 2, 2, 2, 2}, {0, 0, 1, 0, 2, 2, 2, 2}, {0, 0, 1, 1, 0, 0, 0, 0}, {0, 0, 0, 1, 1, 1, 0, 0}, {2, 2, 0, 0, 0, 1, 0, 0}, {2, 2, 2, 0, 0, 1, 0, 0}, {2, 2, 2, 0, 0, 1, 0, 0}}
	result = readFloorFromFile("../floor-files/beaupasbeau")
	if !sont_identique(result_voulu, result) {
		fmt.Println("Vous renvoyer :", result, "alors que le resultat voulu est : ", result_voulu)
		t.Fail()
	}
}

func Test_exemple(t *testing.T) {
	var result [][]tiles.Tiles
	var result_voulu [][]int = [][]int{{1, 1, 3, 4}, {1, 1, 4, 3}, {0, 0, 2, 2}, {0, 0, 2, 2}}
	result = readFloorFromFile("../floor-files/exemple")
	if !sont_identique(result_voulu, result) {
		fmt.Println("Vous renvoyer :", result, "alors que le resultat voulu est : ", result_voulu)
		t.Fail()
	}
}

func Test_exemple2(t *testing.T) {
	var result [][]tiles.Tiles
	var result_voulu [][]int = [][]int{
		{1, 2, 1, 2, 1},
		{1, 2, 2, 1, 2},
		{3, 3, 3, 3, 3},
		{3, 3, 3, 3, 3}}
	result = readFloorFromFile("../floor-files/exemple2")
	if !sont_identique(result_voulu, result) {
		fmt.Println("Vous renvoyer :", result, "alors que le resultat voulu est : ", result_voulu)
		t.Fail()
	}
}

func Test_test_long(t *testing.T) {
	var result [][]tiles.Tiles
	var result_voulu [][]int = [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8},
	}
	result = readFloorFromFile("../floor-files/test_long")
	if !sont_identique(result_voulu, result) {
		fmt.Println("Vous renvoyer :", result, "alors que le resultat voulu est : ", result_voulu)
		t.Fail()
	}
}

func Test_test_haut(t *testing.T) {
	var result [][]tiles.Tiles
	var result_voulu [][]int = [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
		{7, 8},
		{8, 9},
	}
	result = readFloorFromFile("../floor-files/test_haut")
	if !sont_identique(result_voulu, result) {
		fmt.Println("Vous renvoyer :", result, "alors que le resultat voulu est : ", result_voulu)
		t.Fail()
	}
}
