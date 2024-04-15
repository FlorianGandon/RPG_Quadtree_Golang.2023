package particle

func Update(TabParticle [10]*Particle, sens bool) {
	if !sens {
		for i := 0; i < 10; i++ {
			TabParticle[i].Orientation = TabParticle[i].Orientation + 5
			if TabParticle[i].Orientation >= 360 {
				TabParticle[i].Orientation -= 360
			}
		}
	} else {
		for i := 0; i < 10; i++ {
			TabParticle[i].Orientation = TabParticle[i].Orientation - 5
			if TabParticle[i].Orientation < 0 {
				TabParticle[i].Orientation += 360
			}
		}
	}
}
