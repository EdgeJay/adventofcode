package structs

type Submarine struct {
	Horizontal int
	Depth      int
	Aim        int
}

func NewSubmarine() *Submarine {
	return &Submarine{0, 0, 0}
}
