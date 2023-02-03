package kata

/* Complete the function so that it finds the average of the three scores passed to it
and returns the letter value associated with that grade.

Numerical Score		Letter Grade
90 <= score <= 100	'A'
80 <= score < 90	'B'
70 <= score < 80	'C'
60 <= score < 70	'D'
0 <= score < 60		'F'
*/

func GetGrade(a, b, c int) rune {
	avg := (a + b + c) / 3
	var grade rune
	switch {
	case avg >= 90:
		grade = 'A'
	case avg >= 80:
		grade = 'B'
	case avg >= 70:
		grade = 'C'
	case avg >= 60:
		grade = 'D'
	default:
		grade = 'F'
	}
	return grade
}
