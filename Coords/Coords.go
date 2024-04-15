package Coords

import "encoding/json"

type Coords struct {
	X int
	Y int
}

// MarshalJSON Fonction pour encoder le type Coords en JSON (Obligatoire pour la librairie json ¯\_(ツ)_/¯)
func (c Coords) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]int{"X": c.X, "Y": c.Y})
}

// UnmarshalJSON Fonction pour décoder le type Coords depuis JSON (De même)
func (c *Coords) UnmarshalJSON(data []byte) error {
	var tmp map[string]int
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	c.X = tmp["X"]
	c.Y = tmp["Y"]
	return nil
}
