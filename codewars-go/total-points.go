package kata

// match is represented by a string in the format "x:y", ["3:1", "2:2", "0:1", ...]
// Points are awarded for each match as follows:
// if x > y: 3 points (win)
// if x < y: 0 points (loss)
// if x = y: 1 point (tie)
// returns the number of points our team (x) got in the championship

func Points(games []string) int {
	sum := 0
	for _, v := range games {
		if v[0] > v[2] {
			sum += 3
		} else if v[0] == v[2] {
			sum += 1
		}
	}
	return sum
}
