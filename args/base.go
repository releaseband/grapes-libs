package args

type RNG interface {
	Random(min,max uint32) uint32
}