package math

import (
	"github.com/releaseband/grapes-libs/v2/reels"
	"github.com/releaseband/grapes-libs/v2/weights"
)

type Reels interface {
	BuyFeatureReelsGroup() (reels.BuyFeatureReelsGroup, error)
	SimpleReelsGroup() (reels.SimpleReelsGroup, error)
	ExtendedReelsGroup() (reels.ExtendedReelsGroup, error)
	BfReSpinReels() (reels.BuyFeatureReelsGenerator, error)
	BfJackpotReels() (reels.BuyFeatureReelsGenerator, error)
}

type Weights interface {
	Weights() []weights.SimpleWeights
	ExtendedWeights() []weights.ExtendedWeights
}

type Math interface {
	Reels
	Weights
	RTP() uint8
	EXP() float64
}

type Store interface {
	Get(name string) ([]Math, error)
}
