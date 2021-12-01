package reels

import "github.com/releaseband/rng-plugin-adapter/adapter"

type ReelsGenerator interface {
	Generate(rng adapter.RNG) ([][]uint32, error)
}

type ExtendedReels interface {
	Generate(rng adapter.RNG, option uint32) ([][]uint32, error)
}

type BuyFeatureReelsGenerator interface {
	Generate(rng adapter.RNG, bfType uint32, count uint16) ([][]uint32, error)
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
