package reels

type BuyFeatureReelsGroup interface {
	BuyFeatureReels(bfType uint32) ReelsGenerator
	BonusReels() ReelsGenerator
}

type SimpleReelsGroup interface {
	BaseReels() ReelsGenerator
	BonusReels() ReelsGenerator
}
