package reels

type BuyFeatureReelsGroup interface {
	BuyFeatureReels() (BuyFeatureReelsGenerator, error)
	BonusReels() (ReelsGenerator, error)
	ExtendedBonusReels() (ExtendedReels, error)
}

type SimpleReelsGroup interface {
	BaseReels() (ReelsGenerator, error)
	BonusReels() (ReelsGenerator, error)
}

type ExtendedReelsGroup interface {
	SimpleBaseReels() (ReelsGenerator, error)
	BaseReels() (ExtendedReels, error)
	BonusReels() (ExtendedReels, error)
}
