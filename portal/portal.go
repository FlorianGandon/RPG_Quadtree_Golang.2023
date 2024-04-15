package portal

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/particle"

// Contient les deux teleporteur
type Portal struct {
	teleportA *teleport
	teleportB *teleport

	// le dernier téléporteur à être placer
	lastPut int

	// la durrée du message d'erreur
	timeMessage int
}

// X, Y représente les coordonnées en case
// exist indique si le joueur peut voir et intéragir avec ce Teleporteur
type teleport struct {
	X, Y  int
	exist bool

	// le coin de l'image dans la ficher terrain
	topleftImageX int
	topleftImageY int

	particles     [10]*particle.Particle
	sensParticles bool
}
