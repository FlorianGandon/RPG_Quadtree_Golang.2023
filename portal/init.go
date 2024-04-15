package portal

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/particle"
)

func (p *Portal) Init() {
	if configuration.Global.ActiveParticlesPortal {
		p.teleportA = &teleport{exist: false,
			topleftImageX: 0,
			topleftImageY: 64,
			particles:     particle.New(0),
			sensParticles: false}
		p.teleportB = &teleport{exist: false,
			topleftImageX: 16,
			topleftImageY: 64,
			particles:     particle.New(1),
			sensParticles: true}
	} else {
		p.teleportA = &teleport{exist: false,
			topleftImageX: 0,
			topleftImageY: 64}
		p.teleportB = &teleport{exist: false,
			topleftImageX: 16,
			topleftImageY: 64}
	}
}
