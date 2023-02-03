package kata

func GetPlanetName(ID int) string {
	var galaxy = []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune", "Pluto"}
	return galaxy[ID-1]
}
