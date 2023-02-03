package kata

// Let's play! You have to return which player won! In case of a draw return Draw!.
// "scissors", "paper" --> "Player 1 won!"
//"scissors", "rock" --> "Player 2 won!"
// "paper", "paper" --> "Draw!"

func Rps(p1, p2 string) string {
	result := "Player 2 won!"

	m := map[string]string{"rock": "scissors", "paper": "rock", "scissors": "paper"}

	if p1 == p2 {
		result = "Draw!"
	} else if m[p1] == p2 {
		result = "Player 1 won!"
	}
	return result
}
