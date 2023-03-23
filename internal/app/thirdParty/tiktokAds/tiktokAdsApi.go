package tiktokAds

type Campaign struct {
}

type TiktokAdsAPI interface {
	GetCampaigns() ([]Campaign, error)
	CreateCampaign(campaign Campaign) error
	UpdateCampaign(campaign Campaign) error
	DeleteCampaign(campaign Campaign) error
}
