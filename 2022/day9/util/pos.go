package util

import (
	"fmt"

	"github.com/EdgeJay/adventofcode/common/math"
	"github.com/thoas/go-funk"
)

const (
	DIR_UP    = "U"
	DIR_DOWN  = "D"
	DIR_LEFT  = "L"
	DIR_RIGHT = "R"
	DIR_STAY  = ""
)

type Step struct {
	Direction string
	Units     int
}

type VisitedPosition struct {
	X int
	Y int
}

type Position struct {
	Name             string
	X                int
	Y                int
	VisitedPositions []VisitedPosition
}

func NewPosition(name string, x, y int) *Position {
	return &Position{
		Name:             name,
		X:                x,
		Y:                y,
		VisitedPositions: make([]VisitedPosition, 0),
	}
}

func (p *Position) IsAdjacent(target *Position) bool {
	// overlapping
	if p.X == target.X && p.Y == target.Y {
		return true
	}

	// diagonally adjacent
	if (p.X-1 == target.X && p.Y-1 == target.Y) ||
		(p.X-1 == target.X && p.Y+1 == target.Y) ||
		(p.X+1 == target.X && p.Y-1 == target.Y) ||
		(p.X+1 == target.X && p.Y+1 == target.Y) {
		return true
	}

	// next to each other
	if p.Y == target.Y && (p.X+1 == target.X || p.X-1 == target.X) {
		return true
	}

	if p.X == target.X && (p.Y+1 == target.Y || p.Y-1 == target.Y) {
		return true
	}

	return false
}

func (p *Position) Move(step Step, print bool) {
	if print {
		fmt.Printf("%s moved %s %d\n", p.Name, step.Direction, step.Units)
	}

	switch step.Direction {
	case DIR_LEFT:
		p.X -= step.Units
	case DIR_RIGHT:
		p.X += step.Units
	case DIR_UP:
		p.Y += step.Units
	case DIR_DOWN:
		p.Y -= step.Units
	}
}

func (p *Position) MoveNextTo(target *Position) {
	fmt.Printf("%s current position is (%d, %d), move next to (%d, %d)\n", p.Name, p.X, p.Y, target.X, target.Y)

	if !p.IsAdjacent(target) {
		xDiff := math.GetAbsInt(p.X - target.X)
		yDiff := math.GetAbsInt(p.Y - target.Y)

		xDir := DIR_STAY
		if p.X > target.X {
			xDir = DIR_LEFT
		} else if p.X < target.X {
			xDir = DIR_RIGHT
		}

		yDir := DIR_STAY
		if p.Y > target.Y {
			yDir = DIR_DOWN
		} else if p.Y < target.Y {
			yDir = DIR_UP
		}

		// move diagonally
		if xDir != DIR_STAY && yDir != DIR_STAY {
			for {
				if xDiff != 0 {
					p.Move(Step{Direction: xDir, Units: 1}, false)
				}

				if yDiff != 0 {
					p.Move(Step{Direction: yDir, Units: 1}, false)
				}

				p.RecordVisit()

				xDiff--
				yDiff--

				if p.IsAdjacent(target) {
					break
				}
			}
			// or else move along axis
		} else if xDir != DIR_STAY {
			for {
				p.Move(Step{Direction: xDir, Units: 1}, false)
				p.RecordVisit()
				xDiff--

				if p.IsAdjacent(target) {
					break
				}
			}
		} else if yDir != DIR_STAY {
			for {
				p.Move(Step{Direction: yDir, Units: 1}, false)
				p.RecordVisit()
				yDiff--

				if p.IsAdjacent(target) {
					break
				}
			}
		}

		fmt.Printf("total visited positions: %d\n", len(p.VisitedPositions))
	} else {
		fmt.Printf("%s already adjacent to %s\n", p.Name, target.Name)
	}
}

func (p *Position) RecordVisit() {
	pos := VisitedPosition{
		X: p.X,
		Y: p.Y,
	}

	// don't record if visited already
	alreadyVisited := funk.Contains(p.VisitedPositions, pos)

	if !alreadyVisited {
		p.VisitedPositions = append(p.VisitedPositions, pos)
		fmt.Printf("%s recorded position %v\n", p.Name, pos)
	} else {
		fmt.Printf("%s already visited position %v\n", p.Name, pos)
	}
}
