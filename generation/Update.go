package generation

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/Coords"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/tiles"
)

// Update met à jour le quadtree à afficher
func Update(ListChunkgenerate map[Coords.Coords][][]tiles.Tiles) (q quadtree.Quadtree, gx, gy *int) {
	var ListCoordsChunk, minX, maxX, minY, maxY = getListChunkCoordsToShow()
	var toshow = generateifneeded(ListCoordsChunk, ListChunkgenerate)
	gx = &minX
	gy = &minY
	q = Show(toshow, maxX, maxY, minX, minY)
	return
}

// GetCoordsChunkEtremityToShow done les coordonnés des points en haut à gauche et droite et en bas à gauche
func GetCoordsChunkEtremityToShow() (a, b, c Coords.Coords) {
	var x, y = configuration.Global.CameraX, configuration.Global.CameraY
	a = Coords.Coords{x - configuration.Global.ScreenCenterTileX, y - configuration.Global.ScreenCenterTileY}
	b = Coords.Coords{a.X + configuration.Global.NumTileX - 1, a.Y}
	c = Coords.Coords{a.X, a.Y + configuration.Global.NumTileY - 1}
	return
}

// Donne la liste des coordonnés de chunk aà afficher
func getListChunkCoordsToShow() (final []Coords.Coords, minX, maxX, minY, maxY int) {
	var a, b, c = GetCoordsChunkEtremityToShow()
	if a.X >= 0 {
		minX = a.X/configuration.Global.ChunkSize - 1
	} else {
		minX = (a.X-configuration.Global.ChunkSize)/configuration.Global.ChunkSize - 1
	}
	if a.Y >= 0 {
		minY = a.Y/configuration.Global.ChunkSize - 1
	} else {
		minY = (a.Y-configuration.Global.ChunkSize)/configuration.Global.ChunkSize - 1
	}
	maxX, maxY = b.X/configuration.Global.ChunkSize+1, c.Y/configuration.Global.ChunkSize+1
	for cy := minY; cy <= maxY; cy++ {
		for cx := minX; cx <= maxX; cx++ {
			final = append(final, Coords.Coords{cx, cy})
		}
	}
	return
}

// generateifneeded Génère les chunks aux coordonnés fournit si se n'est pas déja fait
func generateifneeded(ListCoordsChunk []Coords.Coords, ListChunkgenerate map[Coords.Coords][][]tiles.Tiles) map[Coords.Coords][][]tiles.Tiles {
	var f = make(map[Coords.Coords][][]tiles.Tiles, 20)
	for _, coords := range ListCoordsChunk {
		if ListChunkgenerate[coords] == nil {
			var h, b, g, d []tiles.Tiles
			var coordToLook = Coords.Coords{coords.X, coords.Y - 1}
			if ListChunkgenerate[coordToLook] != nil {
				h = ListChunkgenerate[coordToLook][len(ListChunkgenerate[coordToLook])-1]
			}
			coordToLook = Coords.Coords{coords.X, coords.Y + 1}
			if ListChunkgenerate[coordToLook] != nil {
				b = ListChunkgenerate[coordToLook][0]
			}
			coordToLook = Coords.Coords{coords.X - 1, coords.Y}
			if ListChunkgenerate[coordToLook] != nil {
				for _, ligne := range ListChunkgenerate[coordToLook] {
					g = append(g, ligne[len(ligne)-1])
				}
			}
			coordToLook = Coords.Coords{coords.X + 1, coords.Y}
			if ListChunkgenerate[coordToLook] != nil {
				for _, ligne := range ListChunkgenerate[coordToLook] {
					d = append(d, ligne[0])
				}
			}
			ListChunkgenerate[coords] = Generation(configuration.Global.ChunkSize, h, b, g, d)
		}
		f[coords] = ListChunkgenerate[coords]
	}
	return f
}

// Show turn chunks to quadtree
func Show(chunkgenerate map[Coords.Coords][][]tiles.Tiles, maxX, maxY, minX, minY int) (q quadtree.Quadtree) {
	var lX = (maxX - minX) + 1
	var lY = (maxY - minY) + 1
	var final = make([][]tiles.Tiles, lY*configuration.Global.ChunkSize)
	for i := range final {
		final[i] = make([]tiles.Tiles, lX*configuration.Global.ChunkSize)
	}
	for coords, chunks := range chunkgenerate {
		for y, c := range chunks {
			for x := range c {
				final[coords.Y*configuration.Global.ChunkSize+y-(minY*configuration.Global.ChunkSize)][coords.X*configuration.Global.ChunkSize+x-(minX*configuration.Global.ChunkSize)] = chunks[y][x]
			}
		}
	}
	q.Width = len(final[0])
	q.Height = len(final)
	q.Root = quadtree.Build_quadtree(final, &q)
	return
}
