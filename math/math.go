package math

import (
	"github.com/releaseband/grapes-libs/reels"
	"github.com/releaseband/grapes-libs/weights"
)

type Math interface {
	RTP() uint8
	EXP() float64
	BuyFeatureReelsGroup() (reels.BuyFeatureReelsGroup, error)
	SimpleReelsGroup() (reels.SimpleReelsGroup, error)
	Weights() []weights.SimpleWeights
}

type Store interface {
	Get(name string) ([]Math, error)
}
