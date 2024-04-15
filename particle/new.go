package particle

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"image/color"
)

func New_Particules(r, g, b uint8, distance float64) Particle {
	return Particle{color.NRGBA{R: r, G: g, B: b, A: 120}, 2, 0, distance}
} //4

func New(color int) (result [10]*Particle) {
	var (
		r uint8
		g uint8
		b uint8
	)

	if color == 0 {
		r = 255
	} else if color == 1 {
		b = 255
	}

	// 4 première sont dans le portail
	var toAdd0 Particle = New_Particules(r/3, g/3, b/3, float64(configuration.Global.TileSize/4))
	var toAdd1 Particle = New_Particules(r/3, g/3, b/3, float64(configuration.Global.TileSize/4))
	var toAdd2 Particle = New_Particules(r/3, g/3, b/3, float64(configuration.Global.TileSize/4))
	var toAdd3 Particle = New_Particules(r/3, g/3, b/3, float64(configuration.Global.TileSize/4))
	// les 6 dernièrs sont à l'extérieur
	var toAdd4 Particle = New_Particules(r, g, b, float64(configuration.Global.TileSize/2))
	var toAdd5 Particle = New_Particules(r, g, b, float64(configuration.Global.TileSize/2))
	var toAdd6 Particle = New_Particules(r, g, b, float64(configuration.Global.TileSize/2))
	var toAdd7 Particle = New_Particules(r, g, b, float64(configuration.Global.TileSize/2))
	var toAdd8 Particle = New_Particules(r, g, b, float64(configuration.Global.TileSize/2))
	var toAdd9 Particle = New_Particules(r, g, b, float64(configuration.Global.TileSize/2))
	result[0] = &toAdd0
	result[1] = &toAdd1
	result[2] = &toAdd2
	result[3] = &toAdd3
	result[4] = &toAdd4
	result[5] = &toAdd5
	result[6] = &toAdd6
	result[7] = &toAdd7
	result[8] = &toAdd8
	result[9] = &toAdd9

	return
}
