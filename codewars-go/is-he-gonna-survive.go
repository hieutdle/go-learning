package kata

// each dragon takes 2 bullets to be defeated
// he's gonna grab a specific given number of bullets
// will he survive?
// Return True if yes, False otherwise :)

func Hero(bullets, dragons int) bool {
	return bullets/2 >= dragons
}
