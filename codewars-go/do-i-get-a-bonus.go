package kata

import "strconv"

// If bonus is true, the salary should be multiplied by 10.
//If bonus is false, the fatcat did not make enough money and must receive only his stated salary.

func BonusTime(salary int, bonus bool) string {
	if bonus {
		salary *= 10
	}
	return "£" + strconv.Itoa(salary)
}
