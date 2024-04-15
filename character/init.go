package character

import (
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/particle"
	"image/color"
)

// Init met en place un personnage. Pour le moment
// cela consiste simplement à initialiser une variable
// responsable de définir l'étape d'animation courante.
func (c *Character) Init() {
	c.animationStep = 1
	var shadow particle.Particle = particle.Particle{Color: color.NRGBA{R: 100, G: 100, B: 100, A: 200}, Radius: 6}
	c.Shadow = &shadow
}
