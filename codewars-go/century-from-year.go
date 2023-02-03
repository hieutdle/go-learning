package kata

/*
Given a year, return the century it is in.
1705 --> 18
1900 --> 19
1601 --> 17
2000 --> 20
*/

func solution(year int) int {
	return (year + 99) / 100
}
