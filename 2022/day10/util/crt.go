package util

import "math"

type CRT struct {
	Cycle          int
	SpritePosition int
	Rows           string
}

func NewCRT(startPosition int) *CRT {
	return &CRT{
		Cycle:          0,
		SpritePosition: startPosition,
		Rows:           "",
	}
}

func (crt *CRT) IsSpriteVisible(drawPosition int) bool {
	return crt.SpritePosition-1 == drawPosition ||
		crt.SpritePosition == drawPosition ||
		crt.SpritePosition+1 == drawPosition
}

func (crt *CRT) Draw() {
	crt.Cycle++

	if math.Mod(float64(crt.Cycle), 40) == 1 {
		crt.Rows += "\n"
	}

	// draw positions are zero-indexed
	drawPosition := int(math.Mod(float64(crt.Cycle), 40) - 1)
	if crt.IsSpriteVisible(drawPosition) {
		crt.Rows += "#"
	} else {
		crt.Rows += "."
	}
}

func (crt *CRT) Noop() {
	crt.Draw()
}

func (crt *CRT) Addx(value int) {
	crt.Draw()

	crt.Draw()
	crt.SpritePosition += value
}

func (crt *CRT) Print() string {
	return crt.Rows
}
