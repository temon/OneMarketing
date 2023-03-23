package tiktokAdsV1_3

type CampaignResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	List     []Campaign `json:"list"`
	PageInfo PageInfo   `json:"page_info"`
}

type Campaign struct {
	SecondaryStatus         string   `json:"secondary_status"`
	Budget                  float64  `json:"budget"`
	DeepBidType             string   `json:"deep_bid_type"`
	AdvertiserID            string   `json:"advertiser_id"`
	RoasBid                 float64  `json:"roas_bid"`
	OperationStatus         string   `json:"operation_status"`
	ObjectiveType           string   `json:"objective_type"`
	IsNewStructure          bool     `json:"is_new_structure"`
	CampaignName            string   `json:"campaign_name"`
	CampaignID              string   `json:"campaign_id"`
	CampaignType            string   `json:"campaign_type"`
	BudgetMode              string   `json:"budget_mode"`
	Objective               string   `json:"objective"`
	ModifyTime              string   `json:"modify_time"`
	CreateTime              string   `json:"create_time"`
	AppPromotionType        string   `json:"app_promotion_type"`
	BudgetOptimizeOn        bool     `json:"budget_optimize_on"`
	BidType                 int      `json:"bid_type"`
	OptimizationGoal        string   `json:"optimization_goal"`
	IsSmartPerformance      bool     `json:"is_smart_performance_campaign"`
	DeepBidTypeDescription  string   `json:"deep_bid_type_description"`
	AppID                   string   `json:"app_id"`
	ROASBidDescription      string   `json:"roas_bid_description"`
	CampaignAppProfileState string   `json:"campaign_app_profile_page_state"`
	SpecialIndustries       []string `json:"special_industries"`
	RFCampaignType          string   `json:"rf_campaign_type"`
}

type PageInfo struct {
	Page        int `json:"page"`
	PageSize    int `json:"page_size"`
	TotalPage   int `json:"total_page"`
	TotalNumber int `json:"total_number"`
}
