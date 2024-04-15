package particle

func Init(TabParticle [10]*Particle) {
	// 4 première sont dans le portail
	TabParticle[0].Orientation = 0
	TabParticle[1].Orientation = 90
	TabParticle[2].Orientation = 180
	TabParticle[3].Orientation = 270

	// les 6 dernièrs sont à l'extérieur
	TabParticle[4].Orientation = 0
	TabParticle[5].Orientation = 60
	TabParticle[6].Orientation = 120
	TabParticle[7].Orientation = 180
	TabParticle[8].Orientation = 240
	TabParticle[9].Orientation = 300
}
