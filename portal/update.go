package portal

import (
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/configuration"
	"github.com/FlorianGandon/RPG_Quadtree_Golang.2023/particle"
)

// renvoie si le joueur peut se téléporter et où
func (p *Portal) CanUseTeleporter(playeurX, playeurY int) (possible bool, newX int, newY int) {
	if p.teleportA.X == playeurX && p.teleportA.Y == playeurY && p.teleportA.exist && p.teleportB.exist {
		return true, p.teleportB.X, p.teleportB.Y
	} else if p.teleportB.X == playeurX && p.teleportB.Y == playeurY && p.teleportA.exist && p.teleportB.exist {
		return true, p.teleportA.X, p.teleportA.Y
	}
	p.timeMessage = 180
	return false, 0, 0
}

// Permet de placer ou enlever un téléporteur
func (p *Portal) Interact(playeurX, playeurY int) {
	if (p.teleportA.X == playeurX && p.teleportA.Y == playeurY && p.teleportA.exist) || (p.teleportB.X == playeurX && p.teleportB.Y == playeurY && p.teleportB.exist) {
		p.remove(playeurX, playeurY)
	} else {
		p.put(playeurX, playeurY)
	}
}

// "place" un téléporteur
func (p *Portal) put(playeurX, playeurY int) {
	// si le téléporteur A n'existe pas
	if !p.teleportA.exist {
		p.teleportA.exist = true
		p.teleportA.X = playeurX
		p.teleportA.Y = playeurY
		if configuration.Global.ActiveParticlesPortal {
			particle.Init(p.teleportA.particles)
		}
		p.lastPut = 1
		// si le téléporteur B n'existe pas
	} else if !p.teleportB.exist {
		p.teleportB.exist = true
		p.teleportB.X = playeurX
		p.teleportB.Y = playeurY
		if configuration.Global.ActiveParticlesPortal {
			particle.Init(p.teleportB.particles)
		}
		p.lastPut = 2
	} else if p.lastPut == 1 {
		p.teleportB.exist = true
		p.teleportB.X = playeurX
		p.teleportB.Y = playeurY
		if configuration.Global.ActiveParticlesPortal {
			particle.Init(p.teleportB.particles)
		}
		p.lastPut = 2
	} else if p.lastPut == 2 {
		p.teleportA.exist = true
		p.teleportA.X = playeurX
		p.teleportA.Y = playeurY
		if configuration.Global.ActiveParticlesPortal {
			particle.Init(p.teleportA.particles)
		}
		p.lastPut = 1
	}
}

// "supprime" un téléporteur
func (p *Portal) remove(playeurX, playeurY int) {
	if p.teleportA.X == playeurX && p.teleportA.Y == playeurY {
		p.teleportA.exist = false
	} else if p.teleportB.X == playeurX && p.teleportB.Y == playeurY {
		p.teleportB.exist = false
	}
}

func (p *Portal) Update() {
	if configuration.Global.ActiveParticlesPortal {
		particle.Update(p.teleportA.particles, p.teleportA.sensParticles)
		particle.Update(p.teleportB.particles, p.teleportB.sensParticles)
	}
}
