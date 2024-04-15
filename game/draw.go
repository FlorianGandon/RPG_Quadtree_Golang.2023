package game

import (
	"fmt"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/generation"
	"image/color"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Draw permet d'afficher à l'écran tous les éléments du jeu
// (le sol, le personnage, les éventuelles informations de debug).
// Il faut faire attention à l'ordre d'affichage pour éviter d'avoir
// des éléments qui en cachent d'autres.
func (g *Game) Draw(screen *ebiten.Image) {
	g.floor.Draw(screen, g.character.XShift, g.character.YShift, g.camera.FluideX, g.camera.FluideY)
	g.portal.Draw(screen, g.camera.X, g.camera.Y, g.character.XShift, g.character.YShift, g.floor.Width, g.floor.Height, g.camera.FluideX, g.camera.FluideY)

	g.character.Draw(screen, g.camera.X, g.camera.Y, g.floor.Width, g.floor.Height, g.camera.FluideX, g.camera.FluideY)

	if configuration.Global.DebugMode {
		g.drawDebug(screen, g.character.XShift, g.character.YShift, g.camera.FluideX, g.camera.FluideY)
	}
}

// drawDebug se charge d'afficher les informations de debug si
// l'utilisateur le demande (positions absolues du personnage
// et de la caméra, grille avec les coordonnées, etc).
func (g Game) drawDebug(screen *ebiten.Image, Xshift, Yshift int, FluideX, FluideY bool) {

	gridColor := color.NRGBA{R: 255, G: 255, B: 255, A: 63}
	gridHoverColor := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	gridLineSize := 2
	chunkColor := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	cameraColor := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	cameraLineSize := 1

	mouseX, mouseY := ebiten.CursorPosition()

	xMaxPos := configuration.Global.ScreenWidth
	yMaxPos := configuration.Global.ScreenHeight

	for x := 0; x < configuration.Global.NumTileX; x++ {
		xGeneralPos := x*configuration.Global.TileSize + configuration.Global.TileSize/2
		xPos := float32(xGeneralPos)
		if FluideX {
			xPos -= float32(Xshift)
		}

		lineColor := gridColor
		if mouseX+1 >= xGeneralPos && mouseX+1 <= xGeneralPos+gridLineSize {
			lineColor = gridHoverColor
		}

		vector.StrokeLine(screen, xPos, 0, xPos, float32(yMaxPos), float32(gridLineSize), lineColor, false)

		xPrintValue := g.camera.X + x - configuration.Global.ScreenCenterTileX
		xPrint := fmt.Sprint(xPrintValue)
		if len(xPrint) <= (2*configuration.Global.TileSize)/16 || (xPrintValue > 0 && xPrintValue%2 == 0) || (xPrintValue < 0 && (-xPrintValue)%2 == 0) {
			xTextPos := xGeneralPos - 3*len(xPrint) - 1
			ebitenutil.DebugPrintAt(screen, xPrint, xTextPos, yMaxPos)
		}
	}

	for y := 0; y < configuration.Global.NumTileY; y++ {
		yGeneralPos := y*configuration.Global.TileSize + configuration.Global.TileSize/2
		yPos := float32(yGeneralPos)
		if FluideY {
			yPos -= float32(Yshift)
		}
		lineColor := gridColor
		if mouseY+1 >= yGeneralPos && mouseY+1 <= yGeneralPos+gridLineSize {
			lineColor = gridHoverColor
		}

		vector.StrokeLine(screen, 0, yPos, float32(xMaxPos), yPos, float32(gridLineSize), lineColor, false)

		yPrint := fmt.Sprint(g.camera.Y + y - configuration.Global.ScreenCenterTileY)
		xTextPos := xMaxPos + 1
		yTextPos := yGeneralPos - 8
		ebitenutil.DebugPrintAt(screen, yPrint, xTextPos, yTextPos)
	}

	if !configuration.Global.BetterCharacter {
		vector.StrokeRect(screen, float32(configuration.Global.ScreenCenterTileX*configuration.Global.TileSize /*+xShift*/), float32(configuration.Global.ScreenCenterTileY*configuration.Global.TileSize /*+yShift*/), float32(configuration.Global.TileSize+1), float32(configuration.Global.TileSize+1), float32(cameraLineSize), cameraColor, false)
	}
	ySpace := 16
	spacing := ySpace
	ebitenutil.DebugPrintAt(screen, "Camera:", xMaxPos+2*configuration.Global.TileSize, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("(", g.camera.X, ",", g.camera.Y, ")"), xMaxPos+2*configuration.Global.TileSize+configuration.Global.TileSize/2, spacing)

	spacing += 2 * ySpace
	ebitenutil.DebugPrintAt(screen, "Character:", xMaxPos+2*configuration.Global.TileSize, spacing)
	spacing += ySpace
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("(", g.character.X, ",", g.character.Y, ")"), xMaxPos+2*configuration.Global.TileSize+configuration.Global.TileSize/2, spacing)

	if configuration.Global.RandomGeneration {
		spacing += 2 * ySpace
		ebitenutil.DebugPrintAt(screen, "Chunk Coordonate:", xMaxPos+2*configuration.Global.TileSize, spacing)
		var chunkX, chunkY int
		if g.camera.X >= 0 {
			chunkX = g.camera.X / configuration.Global.ChunkSize
		} else {
			chunkX = (g.camera.X - configuration.Global.ChunkSize + 1) / configuration.Global.ChunkSize
		}
		if g.camera.Y >= 0 {
			chunkY = g.camera.Y / configuration.Global.ChunkSize
		} else {
			chunkY = (g.camera.Y - configuration.Global.ChunkSize + 1) / configuration.Global.ChunkSize
		}
		spacing += ySpace
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("(", chunkX, ",", chunkY, ")"), xMaxPos+2*configuration.Global.TileSize+configuration.Global.TileSize/2, spacing)

		spacing += 2 * ySpace
		ebitenutil.DebugPrintAt(screen, "Seed:", xMaxPos+2*configuration.Global.TileSize, spacing)
		spacing += ySpace
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("(", configuration.Global.Seed, ")"), xMaxPos+2*configuration.Global.TileSize+configuration.Global.TileSize/2, spacing)
		spacing += ySpace
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("(", *configuration.Global.SeedRand, ")"), xMaxPos+2*configuration.Global.TileSize+configuration.Global.TileSize/2, spacing)
		var a, b, c = generation.GetCoordsChunkEtremityToShow()
		for x := a.X + 1; x <= b.X; x++ {
			if x%configuration.Global.ChunkSize == 0 {
				var xPos = float32(x*configuration.Global.TileSize + configuration.Global.ScreenWidth/2 - configuration.Global.TileSize/2 - configuration.Global.CameraX*configuration.Global.TileSize)
				if FluideX {
					xPos = xPos - float32(Xshift)
				}
				vector.StrokeLine(screen, xPos, 0, xPos, float32(yMaxPos), float32(2), chunkColor, false)
			}
		}
		for y := a.Y + 1; y <= c.Y; y++ {
			if y%configuration.Global.ChunkSize == 0 {
				var yPos = float32(y*configuration.Global.TileSize + configuration.Global.ScreenWidth/2 - configuration.Global.TileSize/2 - configuration.Global.CameraY*configuration.Global.TileSize)
				if FluideY {
					yPos = yPos - float32(Yshift)
				}
				vector.StrokeLine(screen, 0, yPos, float32(xMaxPos), yPos, float32(2), chunkColor, false)
			}
		}
	}
}
