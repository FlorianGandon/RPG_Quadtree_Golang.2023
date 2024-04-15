package character

func (c *Character) Teleport(newX, newY int) {
	c.X = newX
	c.Y = newY
}
