package reels

// deprecated
type BuyFeatureReelsGroup interface {
	BuyFeatureReels() (BuyFeatureReelsGenerator, error)
	BonusReels() (ReelsGenerator, error)
	ExtendedBonusReels() (ExtendedReels, error)
}

// deprecated
type SimpleReelsGroup interface {
	BaseReels() (ReelsGenerator, error)
	BonusReels() (ReelsGenerator, error)
}

// deprecated
type ExtendedReelsGroup interface {
	SimpleBaseReels() (ReelsGenerator, error)
	BaseReels() (ExtendedReels, error)
	BonusReels() (ExtendedReels, error)
}
