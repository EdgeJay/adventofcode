package day4

import (
	"fmt"
	"slices"
)

type CardPile struct {
	input []string
	Cards []*Card
}

func NewCardPile(input []string) *CardPile {
	return &CardPile{
		input: input,
		Cards: nil,
	}
}

func (p *CardPile) findCard(id int) *Card {
	idx := slices.IndexFunc(p.Cards, func(c *Card) bool {
		return c.Id == id
	})

	if idx != -1 {
		return p.Cards[idx]
	}

	return nil
}

func (p *CardPile) traverse(card *Card, total *int) {

	fmt.Println(card.Id, *total)
	count, err := card.NumberOfCardsWon()
	if err != nil {
		return
	}

	*total += count

	for idx := card.Id + 1; idx <= card.Id+count; idx++ {
		nextCard := p.findCard(idx)
		if nextCard != nil {
			p.traverse(nextCard, total)
		}
	}
}

func (p *CardPile) TotalWinningCards() int {

	// Create cards
	cards := make([]*Card, 0)
	for _, line := range p.input {
		card := NewCard(line)
		if err := card.parseInput(); err == nil {
			cards = append(cards, card)
		}
	}

	p.Cards = cards

	total := len(p.Cards)

	for _, card := range p.Cards {
		p.traverse(card, &total)
	}

	return total
}
