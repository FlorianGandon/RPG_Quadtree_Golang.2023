package generation

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/Coords"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/tiles"
)

// Init Première génération au démarrage
func Init() (q quadtree.Quadtree, f map[Coords.Coords][][]tiles.Tiles) {
	var ListChunkgenerate = make(map[Coords.Coords][][]tiles.Tiles)
	var ListCoordsChunk, minX, maxX, minY, maxY = getListChunkCoordsToShow()
	ListChunkgenerate = generateifneeded(ListCoordsChunk, ListChunkgenerate)
	ListChunkgenerate[Coords.Coords{X: 0, Y: 0}][0][0] = tiles.Tiles{Types: 6, Entropy: 0, ListPossibility: nil}
	return Show(ListChunkgenerate, maxX, maxY, minX, minY), ListChunkgenerate
}
