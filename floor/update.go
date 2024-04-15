package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update se charge de stocker dans la structure interne (un tableau)
// de f une représentation de la partie visible du terrain à partir
// des coordonnées absolues de la case sur laquelle se situe la
// caméra.
//
// On aurait pu se passer de cette fonction et tout faire dans Draw.
// Mais cela permet de découpler le calcul de l'affichage.
func (f *Floor) Update(camXPos, camYPos int) {
	//fmt.Println(configuration.Global.ScreenCenterTileX)
	f.water_clock++
	f.water_clock = f.water_clock % 30
	switch configuration.Global.FloorKind {
	case gridFloor:
		f.updateGridFloor(camXPos, camYPos)
	case fromFileFloor:
		f.updateFromFileFloor(camXPos, camYPos)
	case quadTreeFloor:
		f.updateQuadtreeFloor(camXPos, camYPos)
	}
}

// le sol est un quadrillage de tuiles d'herbe et de tuiles de désert
func (f *Floor) updateGridFloor(camXPos, camYPos int) {
	for y := 0; y < len(f.Content); y++ {
		for x := 0; x < len(f.Content[y]); x++ {
			absCamX := camXPos
			if absCamX < 0 {
				absCamX = -absCamX
			}
			absCamY := camYPos
			if absCamY < 0 {
				absCamY = -absCamY
			}
			f.Content[y][x].Types = ((x + absCamX%2) + (y + absCamY%2)) % 2
		}
	}
}

// le sol est récupéré depuis un tableau, qui a été lu dans un fichier
func (f *Floor) updateFromFileFloor(camXPos, camYPos int) {
	// on récupere les coordonnée de la case en haut à gauche, avec un décalage de -1 car on charge une carte plus grande que ce qui doit être afficher pour pouvoir avoir une caméra fluide
	var topLeftX int = camXPos - configuration.Global.ScreenCenterTileX - 1
	var topLeftY int = camYPos - configuration.Global.ScreenCenterTileY - 1

	// on parcours f.content
	for y := 0; y < len(f.Content); y++ {
		for x := 0; x < len(f.Content[0]); x++ {
			// si la tuile qu'on regarde est en dehors du fullContent alors :
			if topLeftX+x < 0 || topLeftY+y < 0 || topLeftX+x >= f.Width || topLeftY+y >= f.Height {
				if configuration.Global.RoundEarth {
					// si le monde est une terre ronde on répète le fullContent
					f.Content[y][x].Types = f.fullContent[(topLeftY+y+(f.Height*(configuration.Global.NumTileY/f.Height)))%f.Height][(topLeftX+x+f.Width*(configuration.Global.NumTileX/f.Width))%+f.Width].Types
				} else {
					// on met la case à -1 sinon
					f.Content[y][x].Types = -1
				}
			} else {
				// sinon on affecte à la case la valeur voulu
				f.Content[y][x].Types = f.fullContent[topLeftY+y][topLeftX+x].Types
			}
		}
	}

}

// le sol est récupéré depuis un quadtree, qui a été lu dans un fichier
func (f *Floor) updateQuadtreeFloor(camXPos, camYPos int) {
	var topLeftX, topLeftY int
	//fmt.Println(f.GX, f.GY)
	if f.GX != nil {
		topLeftX = camXPos - configuration.Global.ScreenCenterTileX - *f.GX*configuration.Global.ChunkSize
	} else {
		topLeftX = camXPos - configuration.Global.ScreenCenterTileX
	}
	if f.GY != nil {
		topLeftY = camYPos - configuration.Global.ScreenCenterTileY - *f.GY*configuration.Global.ChunkSize
	} else {
		topLeftY = camYPos - configuration.Global.ScreenCenterTileY
	}
	//topLeftX += 1
	var width int = len(f.Content[0])
	var height int = len(f.Content)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			f.Content[y][x].Types = -1
		}
	}

	if configuration.Global.RoundEarth {
		for x := 0; x <= (configuration.Global.NumTileX/2)+1; x++ {
			for y := 0; y <= (configuration.Global.NumTileY/2)+1; y++ {
				f.QuadtreeContent.GetContent(f.Width-(x*f.Width)+camXPos+1, f.Height-(y*f.Height)+camYPos+1, f.Content)
			}
		}
	} else {
		f.QuadtreeContent.GetContent(topLeftX, topLeftY, f.Content)
	}
}
