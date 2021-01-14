package bowling_game

type Game struct {
	rolls       [21]int
	currentRoll int
}

func (g *Game) Roll(pins int) {
	g.rolls[g.currentRoll] = pins
	g.currentRoll++
}

func (g *Game) Score() int {
	score := 0
	i := 0

	for frame := 0; frame < 10; frame++ {
		switch {
		case g.isStrike(i):
			score += 10 + g.strikeBonus(i)
			i += 1
		case g.isSpare(i):
			score += 10 + g.spareBonus(i)
			i += 2
		default:
			score += g.rolls[i] + g.rolls[i+1]
			i += 2
		}
	}

	return score
}

func (g *Game) isSpare(firstBallIndex int) bool {
	return g.rolls[firstBallIndex]+g.rolls[firstBallIndex+1] == 10
}

func (g *Game) isStrike(firstBallIndex int) bool {
	return g.rolls[firstBallIndex] == 10
}

func (g *Game) strikeBonus(firstBallIndex int) int {
	return g.rolls[firstBallIndex+1] + g.rolls[firstBallIndex+2]
}

func (g *Game) spareBonus(firstBallIndex int) int {
	return g.rolls[firstBallIndex+2]
}
