package structs

type Submarine struct {
	Horizontal int
	Depth      int
}

func NewSubmarine() *Submarine {
	return &Submarine{0, 0}
}
