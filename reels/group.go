package reels

type BuyFeatureReelsGroup interface {
	BuyFeatureReels() BuyFeatureReelsGenerator
	BonusReels() ReelsGenerator
}

type SimpleReelsGroup interface {
	BaseReels() ReelsGenerator
	BonusReels() ReelsGenerator
}