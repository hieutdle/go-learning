package kata

// If one of the flowers has an even number of petals
// and the other has an odd number of petals it means they are in love.
// return true if they are in love and false if they aren't.

func LoveFunc(flower1, flower2 int) bool {
	return (flower1+flower2)%2 == 1
}
