package quadtree

import (
	"testing"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/tiles"
)

func createContent(n int) (fullContent [][]tiles.Tiles) {
	fullContent = make([][]tiles.Tiles, n)
	for y := 0; y < n; y++ {
		fullContent[y] = make([]tiles.Tiles, n)
	}
	return fullContent
}

func Benchmark(b *testing.B) {
	var fullContent [][]tiles.Tiles = createfullContent(1000)
	var content [][]tiles.Tiles = createContent(1000)
	var q Quadtree = MakeFromArray(fullContent)
	for i := 0; i < b.N; i++ {
		q.GetContent(0, 0, content)
	}
}
