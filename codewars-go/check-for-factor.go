package kata

// 2 and 3 are factors of 6 because: 2 * 3 = 6
// This function should test if the factor is a factor of base.

func CheckForFactor(base int, factor int) bool {
	return base%factor == 0
}
