package eneryDay

type LosingPlayer struct{}

func (losing LosingPlayer) LosingPlayer(x int, y int) string {
	userName := []string{"Alice", "Bob"}
	index := 0
	for x > 0 && y > 3 {
		ySum := (115 - 75) / 10
		y -= ySum
		x--
		index++
	}
	winner := 0
	if index%2 == 1 {
		winner = 0
	} else {
		winner = 1
	}
	return userName[winner]

}
