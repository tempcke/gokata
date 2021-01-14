package bowling_game

import (
	"gotest.tools/assert"
	"testing"
)

// RULES
// Gutter game scores zero - When you roll all misses, you get a total score of zero.
// All ones scores 20 - When you knock down one pin with each ball, your total score is 20.
// A spare in the first frame, followed by three pins, followed by all misses scores 16.
// A strike in the first frame, followed by three and then four pins, followed by all misses, scores 24.
// A perfect game (12 strikes) scores 300.

func TestGutterGame(t *testing.T) {
	g := Game{}
	rollMany(&g, 0, 20)
	assert.Equal(t, 0, g.Score())
}

func TestAllOnes(t *testing.T) {
	g := Game{}
	rollMany(&g, 1, 20)
	assert.Equal(t, 20, g.Score())
}

func TestSpare(t *testing.T) {
	g := Game{}
	g.Roll(5)
	g.Roll(5)
	g.Roll(3)
	rollMany(&g, 17, 0)
	assert.Equal(t, 16, g.Score())
}

func TestStrike(t *testing.T) {
	g := Game{}
	g.Roll(10)
	g.Roll(3)
	g.Roll(4)
	rollMany(&g, 16, 0)
	assert.Equal(t, 24, g.Score())
}

func TestPerfectGame(t *testing.T) {
	g := Game{}
	rollMany(&g, 10, 12)
	assert.Equal(t, 300, g.Score())
}

func rollMany(g *Game, pins int, roles int) {
	for i := 0; i < roles; i++ {
		g.Roll(pins)
	}
}
