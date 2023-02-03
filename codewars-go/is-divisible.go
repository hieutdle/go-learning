package kata

//  checks if a number n is divisible by two numbers x AND y
//  n =  12, x = 2, y = 6 =>  true because  12 is divisible by 2 and 6
//  n = 100, x = 5, y = 3 => false because 100 is not divisible by 3

func IsDivisible(n, x, y int) bool {
	return n%x+n%y == 0
}
