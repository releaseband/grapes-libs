package weights

type Weights interface {
	Value(i uint32) (uint16, bool)
	Gap(value uint16) (uint32, uint32, bool)
	Max() uint32
}

type WeightGetter interface {
	Get(belongs, weightName string, groupKey uint16) Weights
}


// [min,max]
type rng func(min uint32, max uint32) uint32

type GroupedWeights interface {
	Value(groupName string, rng rng) (uint16, error)
}

type GroupedWeightsGetter interface {
	Get(belongs, weightName string, groupKey uint16) GroupedWeights
}