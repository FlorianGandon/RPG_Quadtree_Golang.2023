package particle

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"math"
)

func Draw(screen *ebiten.Image, TabParticle [10]*Particle, posX, posY float64) {
	var sin float64
	var cos float64
	for i := 0; i < 10; i++ {
		sin = math.Sin((TabParticle[i].Orientation * math.Pi) / 180)
		cos = math.Cos((TabParticle[i].Orientation * math.Pi) / 180)
		vector.DrawFilledCircle(screen, float32(posX+8+(cos*TabParticle[i].distance)), float32(posY+8+(sin*TabParticle[i].distance)), TabParticle[i].Radius, TabParticle[i].Color, true)

	}

}

func (p Particle) DrawShadow(screen *ebiten.Image, posX, posY int) {
	vector.DrawFilledCircle(screen, float32(posX+configuration.Global.TileSize/2), float32(posY+12), p.Radius, p.Color, true)
}
