package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetParametersGameId(t *testing.T) {
	game := NewGame()
	game.SetParameters("Game 123: 456 red, 10 green, 5 blue")
	assert.Equal(t, 123, game.Id)
}

func TestSetParametersColours(t *testing.T) {
	game := NewGame()
	game.SetParameters("Game 12: 1 red, 10 blue, 5 green; 11 blue, 6 green; 6 green; 1 green, 1 red, 12 blue; 3 blue; 3 blue, 4 green, 1 red")
	assert.Equal(t, 12, game.Id)
	assert.Equal(t, 1, game.Red)
	assert.Equal(t, 6, game.Green)
	assert.Equal(t, 12, game.Blue)
}

func TestGameWithinLimits(t *testing.T) {
	game := NewGame()
	game.SetParameters("Game 12: 1 red, 10 blue, 5 green; 11 blue, 6 green; 6 green; 1 green, 1 red, 12 blue; 3 blue; 3 blue, 4 green, 1 red")
	assert.Equal(t, true, game.WithinLimits(1, 6, 12))
	assert.Equal(t, false, game.WithinLimits(1, 5, 10))
	assert.Equal(t, false, game.WithinLimits(0, 7, 20))
}
