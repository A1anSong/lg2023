package nonmigrate

type GeoDistributionData struct {
	Day   []GeoDistributionDataItem `json:"day" form:"day"`
	Week  []GeoDistributionDataItem `json:"week" form:"week"`
	Month []GeoDistributionDataItem `json:"month" form:"month"`
	Total []GeoDistributionDataItem `json:"total" form:"total"`
}

type GeoDistributionDataItem struct {
	City  *string `json:"city" form:"city"`
	Count *int    `json:"count" form:"count"`
}
