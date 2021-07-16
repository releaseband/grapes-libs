package reels

import "github.com/releaseband/grapes-libs/rng"

type ReelsGenerator interface {
	GenerateSymbols(rng rng.RNG) ([][]uint32, error)
}

type BuyFeatureReelsGenerator interface {
	GenerateSymbols(rng rng.RNG, bfType uint32) ([][]uint32, error)
}

//deprecated
type Generator interface {
	Type() uint16
	Generate(ids []int) [][]uint32
}

//deprecated
type Group interface {
	Exp() float64
	Generators() []Generator
}

//deprecated
type Getter interface {
	Get(name string, key uint8) Group
}
