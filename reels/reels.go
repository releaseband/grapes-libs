package reels

type RNG interface {
	Random(min,max uint32) uint32
}

type Type interface {
	Type() uint16
}

type ReelsGenerator interface {
	Type
	Generate(rng RNG) ([][]uint32, error)
}

type Generator interface {
	Type
	Generate(ids []int) [][]uint32
}

type Group interface {
	Exp() float64
	Generators() []Generator
}

type Getter interface {
	Get(name string, key uint8) Group
}