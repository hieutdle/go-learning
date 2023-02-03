package kata

// Clock shows h hours, m minutes and s seconds after midnight.
// returns the time since midnight in milliseconds.

func Past(h, m, s int) int {
	result := (((h*60)+m)*60 + s) * 1000
	return result

}
