package tiktokAds

type TiktokAdsAPIAdapter struct {
	Api TiktokAdsAPI
}

func (adapter *TiktokAdsAPIAdapter) GetCampaigns() ([]Campaign, error) {
	return adapter.Api.GetCampaigns()
}

func (adapter *TiktokAdsAPIAdapter) CreateCampaign(campaign Campaign) error {
	return adapter.Api.CreateCampaign(campaign)
}

func (adapter *TiktokAdsAPIAdapter) UpdateCampaign(campaign Campaign) error {
	return adapter.Api.UpdateCampaign(campaign)
}

func (adapter *TiktokAdsAPIAdapter) DeleteCampaign(campaign Campaign) error {
	return adapter.Api.DeleteCampaign(campaign)
}
