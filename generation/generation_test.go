package generation

import (
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/configuration"
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/tiles"
	"testing"
)

func Benchmark(b *testing.B) {
	configuration.Global.ChunkSize = 256
	for i := 0; i < b.N; i++ {
		Generation(256, []tiles.Tiles{}, []tiles.Tiles{}, []tiles.Tiles{}, []tiles.Tiles{})
	}
}
