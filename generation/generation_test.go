package generation

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/tiles"
	"testing"
)

func Benchmark(b *testing.B) {
	configuration.Global.ChunkSize = 256
	for i := 0; i < b.N; i++ {
		Generation(256, []tiles.Tiles{}, []tiles.Tiles{}, []tiles.Tiles{}, []tiles.Tiles{})
	}
}
