package weights

import "github.com/releaseband/rng/v4"

// deprecated
type Weights interface {
	Value(i uint32) (uint16, bool)
	Gap(value uint16) (uint32, uint32, bool)
	Max() uint32
}

// deprecated
type WeightGetter interface {
	Get(belongs, weightName string, groupKey uint16) Weights
}

// deprecated
type GroupedWeights interface {
	Value(groupName string, rng func(uint32, uint32) uint32) (uint16, error)
}

// deprecated
type GroupedWeightsGetter interface {
	Get(belongs, weightName string, groupKey uint16) GroupedWeights
}

// deprecated
type SimpleWeights interface {
	Name() string
	Value(rng rng.RNG) (uint16, error)
}

// deprecated
type ExtendedWeights interface {
	Name() string
	Value(rng rng.RNG, option uint32) (uint16, error)
}
