package kata

// Calculate how many years ago the father was twice as old as his son
// (or in how many years he will be twice as old).
// The answer is always greater or equal to 0

func TwiceAsOld(dadYearsOld int, sonYearsOld int) int {
	x := dadYearsOld - sonYearsOld*2
	if x < 0 {
		x = -x
	}
	return x
}
