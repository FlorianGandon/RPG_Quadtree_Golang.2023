package generation

import (
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/Coords"
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/quadtree"
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/tiles"
)

// Init Première génération au démarrage
func Init() (q quadtree.Quadtree, f map[Coords.Coords][][]tiles.Tiles) {
	var ListChunkgenerate = make(map[Coords.Coords][][]tiles.Tiles)
	var ListCoordsChunk, minX, maxX, minY, maxY = getListChunkCoordsToShow()
	ListChunkgenerate = generateifneeded(ListCoordsChunk, ListChunkgenerate)
	ListChunkgenerate[Coords.Coords{X: 0, Y: 0}][0][0] = tiles.Tiles{Types: 6, Entropy: 0, ListPossibility: nil}
	return Show(ListChunkgenerate, maxX, maxY, minX, minY), ListChunkgenerate
}
