package reels

type BuyFeatureReelsGroup interface {
	BuyFeatureReels(bfType uint8) ReelsGenerator
	BonusReels() ReelsGenerator
}

type SimpleReelsGroup interface {
	BaseReels() ReelsGenerator
	BonusReels() ReelsGenerator
}
