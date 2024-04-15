package quadtree

import (
	"testing"

	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/tiles"
)

func createfullContent(n int) (fullContent [][]tiles.Tiles) {
	fullContent = make([][]tiles.Tiles, n)
	for y := 0; y < n; y++ {
		fullContent[y] = make([]tiles.Tiles, n)
		for x := 0; x < n; x++ {
			fullContent[y][x] = tiles.Tiles{Types: y%2 + x%2}
		}
	}
	return fullContent
}

func BenchmarkMake(b *testing.B) {
	var fullContent [][]tiles.Tiles = createfullContent(1000)
	var q Quadtree
	for i := 0; i < b.N; i++ {
		q = Quadtree{}
		q = MakeFromArray(fullContent)
	}
	q.Root = nil
}

// commandes :
// go tool pprof quadtree.test monprofile.prof
// go test -cpuprofile=monprofile.prof -bench=. -run=quadtreeMakeBench_test.go
// go test -cpuprofile=cpu.prof  -memprofile mem.prof -bench=. -run=quadtreeMakeBench_test.go
