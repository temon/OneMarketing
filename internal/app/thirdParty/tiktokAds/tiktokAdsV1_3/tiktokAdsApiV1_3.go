package tiktokAdsV1_3

import "oneMarketing/internal/app/thirdParty/tiktokAds"

type TiktokAdsAPIv1_3 struct {
	AccessToken string
	BaseURL     string
}

func (api *TiktokAdsAPIv1_3) GetCampaigns() ([]tiktokAds.Campaign, error) {
	// implementation specific to v1.3 of the API
	return nil, nil
}

func (api *TiktokAdsAPIv1_3) CreateCampaign(campaign tiktokAds.Campaign) error {
	// implementation specific to v1.3 of the API
	return nil
}

func (api *TiktokAdsAPIv1_3) UpdateCampaign(campaign tiktokAds.Campaign) error {
	// implementation specific to v1.3 of the API
	return nil
}

func (api *TiktokAdsAPIv1_3) DeleteCampaign(campaign tiktokAds.Campaign) error {
	// implementation specific to v1.3 of the API
	return nil
}
