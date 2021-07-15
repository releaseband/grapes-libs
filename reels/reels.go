package reels

import "github.com/releaseband/grapes-libs/base"

type ReelsGenerator interface {
	Generate(rng base.RNG) ([][]uint32, error)
}

type BuyFeatureReelsGenerator interface {
	Generate(rng base.RNG, bfType uint32) ([][]uint32, error)
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