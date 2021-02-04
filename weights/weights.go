package weights

type Weights interface {
	Value(i uint32) (uint16, bool)
	Gap(value uint16) (uint32, uint32, bool)
	Max() uint32
}

type WeightGetter interface {
	Get(belongs, weightName string, groupKey uint16) Weights
}