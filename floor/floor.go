package floor

import (
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/Coords"
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/quadtree"
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/tiles"
)

// Floor représente les données du terrain. Pour le moment
// aucun champs n'est exporté.
//
//   - content : partie du terrain qui doit être affichée à l'écran
//   - fullContent : totalité du terrain (utilisé seulement avec le type
//     d'affichage du terrain "fromFileFloor")
//   - quadTreeContent : totalité du terrain sous forme de quadtree (utilisé
//     avec le type d'affichage du terrain "quadtreeFloor")
type Floor struct {
	Content           [][]tiles.Tiles
	fullContent       [][]tiles.Tiles
	QuadtreeContent   quadtree.Quadtree
	ListChunkgenerate map[Coords.Coords][][]tiles.Tiles

	water_clock int
	Height      int
	Width       int
	GX          *int
	GY          *int
}

// Types d'affichage du terrain disponibles
const (
	gridFloor int = iota
	fromFileFloor
	quadTreeFloor
)
