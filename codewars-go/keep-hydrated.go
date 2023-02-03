package kata

// Nathan drinks 0.5 litres of water per hour of cycling.
// return the number of litres he will drink
// time = 3 ----> litres = 1
// time = 6.7---> litres = 3
// time = 11.8--> litres = 5

func Litres(time float64) int {
	return int(time / 2)
}
