package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardCalculate(t *testing.T) {

	line := `Card   1: 73 92 13 35 18 96 37 72 76 39 | 82 14 66 57 25 98 49 28  3 95 81 85 31 30 16 79  7 12 55 19 97 45  9 58  2`
	card := NewCard(line)
	_, err := card.Calculate()
	assert.NoError(t, err)
	assert.Equal(t, 1, card.Id)

	assert.Equal(t, 10, len(card.WinningNumbers))
	assert.Equal(t, 73, card.WinningNumbers[0])
	assert.Equal(t, 18, card.WinningNumbers[4])

	assert.Equal(t, 25, len(card.CardNumbers))
	assert.Equal(t, 82, card.CardNumbers[0])
	assert.Equal(t, 3, card.CardNumbers[8])
}

func TestCardCalculateResult(t *testing.T) {

	line := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	card := NewCard(line)
	result, err := card.Calculate()
	assert.NoError(t, err)
	assert.Equal(t, 8, result)
}
