package reels

type BuyFeatureReelsGroup interface {
	BuyFeatureReels() BuyFeatureReelsGenerator
	BonusReels() ReelsGenerator
	Exp() float64
	GroupKey() uint8
}

type SimpleReelsGroup interface {
	BaseReels() ReelsGenerator
	BonusReels() ReelsGenerator
	Exp() float64
	GroupKey() uint8
}