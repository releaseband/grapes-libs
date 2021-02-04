package reels

type Generator interface {
	Generate(ids []int) [][]uint32
	Type() uint16
}

type Group interface {
	Exp() float64
	Generators() []Generator
}

type Getter interface {
	Get(name string, key uint8) Group
}