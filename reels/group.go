package reels

type BuyFeatureReelsGroup interface {
	BuyFeatureReels() (BuyFeatureReelsGenerator, error)
	BonusReels() (ReelsGenerator, error)
}

type SimpleReelsGroup interface {
	BaseReels() (ReelsGenerator, error)
	BonusReels() (ReelsGenerator, error)
}
