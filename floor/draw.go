package floor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/tiles"
	"image"
)

// Draw affiche dans une image (en général, celle qui représente l'écran),
// la partie du sol qui est visible (qui doit avoir été calculée avec Get avant).
func (f Floor) Draw(screen *ebiten.Image, Xshift, Yshift int, FluideX, FluideY bool) {

	for y := range f.Content {
		for x := range f.Content[y] {
			// on parcours le contenue de content pour l'afficher
			if f.Content[y][x].Types != -1 {
				var op *ebiten.DrawImageOptions = &ebiten.DrawImageOptions{}
				var posX float64
				var posY float64

				if FluideX {
					posX = float64((x-1)*configuration.Global.TileSize - Xshift)
				} else {
					posX = float64((x - 1) * configuration.Global.TileSize)
				}
				if FluideY {
					posY = float64((y-1)*configuration.Global.TileSize - Yshift)
				} else {
					posY = float64((y - 1) * configuration.Global.TileSize)
				}
				op.GeoM.Translate(posX, posY)

				shiftX := f.Content[y][x].Types*configuration.Global.TileSize*4 + 16
				var shifty int = 16
				if f.Content[y][x].Types != 0 && configuration.Global.ConnectedTile {
					// terrain plus "naturel"
					var h, b, g, d = Voisins(f.Content, y, x)
					var tileSize = configuration.Global.TileSize
					switch {
					case !h && !b && !g && !d:
						shiftX += tileSize * 2
						shifty += tileSize * 2
					case !h && !b && !g && d:
						shiftX -= tileSize
						shifty += tileSize * 2
					case !h && !b && g && !d:
						shiftX += tileSize
						shifty += tileSize * 2
					case !h && b && !g && !d:
						shiftX += tileSize * 2
						shifty -= tileSize
					case h && !b && !g && !d:
						shiftX += tileSize * 2
						shifty += tileSize

					case !h && !b && g && d:
						shifty += tileSize * 2
					case h && b && !g && !d:
						shiftX += tileSize * 2
					case h && !b && !g && d:
						shiftX -= tileSize
						shifty += tileSize
					case h && !b && g && !d:
						shiftX += tileSize
						shifty += tileSize
					case !h && b && !g && d:
						shiftX -= tileSize
						shifty -= tileSize
					case !h && b && g && !d:
						shiftX += tileSize
						shifty -= tileSize

					case h && b && !g && d:
						shiftX -= tileSize
					case h && b && g && !d:
						shiftX += tileSize
					case !h && b && g && d:
						shifty -= tileSize
					case h && !b && g && d:
						shifty += tileSize
					}

				}

				if f.Content[y][x].Types == 5 {
					shifty = (f.water_clock/10)*configuration.Global.TileSize*4 + shifty
				}
				if float64(configuration.Global.TileSize*configuration.Global.NumTileX) > posX && float64(configuration.Global.TileSize*configuration.Global.NumTileY) > posY {
					// on affiche les tuiles seulement si leurs positions en haut à gauche est dans le cadre
					if float64(configuration.Global.TileSize*configuration.Global.NumTileX) < posX+float64(configuration.Global.TileSize) && float64(configuration.Global.TileSize*configuration.Global.NumTileY) < posY+float64(configuration.Global.TileSize) {
						//cas où le coin en bas à droite n'est pas dans l'écran
						screen.DrawImage(assets.FloorImage.SubImage(
							image.Rect(shiftX, shifty, shiftX+configuration.Global.TileSize-((int(posX)+configuration.Global.TileSize)-(configuration.Global.TileSize*configuration.Global.NumTileX)), shifty+configuration.Global.TileSize-((int(posY)+configuration.Global.TileSize)-(configuration.Global.TileSize*configuration.Global.NumTileY))),
						).(*ebiten.Image), op)
					} else if float64(configuration.Global.TileSize*configuration.Global.NumTileY) < posY+float64(configuration.Global.TileSize) {
						// cas où le coin en bas à gauche n'est pas dans l'écran
						screen.DrawImage(assets.FloorImage.SubImage(
							image.Rect(shiftX, shifty, shiftX+configuration.Global.TileSize, shifty+configuration.Global.TileSize-((int(posY)+configuration.Global.TileSize)-(configuration.Global.TileSize*configuration.Global.NumTileY))),
						).(*ebiten.Image), op)
					} else if float64(configuration.Global.TileSize*configuration.Global.NumTileX) < posX+float64(configuration.Global.TileSize) {
						// cas où le coin en haut à droite n'est pas dans l'écran
						screen.DrawImage(assets.FloorImage.SubImage(
							image.Rect(shiftX, shifty, shiftX+configuration.Global.TileSize-((int(posX)+configuration.Global.TileSize)-(configuration.Global.TileSize*configuration.Global.NumTileX)), shifty+configuration.Global.TileSize),
						).(*ebiten.Image), op)
					} else {
						// affiche normale
						screen.DrawImage(assets.FloorImage.SubImage(
							image.Rect(shiftX, shifty, shiftX+configuration.Global.TileSize, shifty+configuration.Global.TileSize),
						).(*ebiten.Image), op)
					}
				}
			}
		}
	}
}

// renvois les Voisins haut, bas, gauche, droite de la case en coordoné x et y
func Voisins(content [][]tiles.Tiles, y, x int) (h, b, g, d bool) {
	var val = content[y][x]
	if y == 0 {
		h = true
	} else {
		h = seColle(content[y-1][x].Types, val.Types)
	}

	if y == len(content)-1 {
		b = true
	} else {
		b = seColle(content[y+1][x].Types, val.Types)
	}

	if x-1 < 0 {
		g = true
	} else {
		g = seColle(content[y][x-1].Types, val.Types)
	}

	if x+1 >= len(content[y]) {
		d = true
	} else {
		d = seColle(content[y][x+1].Types, val.Types)
	}
	return
}

func seColle(id_casse_cote int, id_case int) bool {
	if id_casse_cote == -1 {
		return true
	}
	var colle [][]int = [][]int{{0, 6}, {1}, {2, 4}, {3}, {4}, {5}, {0, 6}, {7}}
	for i := range colle[id_case] {
		if colle[id_case][i] == id_casse_cote {
			return true
		}
	}
	return false
}
