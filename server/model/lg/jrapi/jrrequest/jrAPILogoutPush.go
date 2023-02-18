package jrrequest

type JRAPILogoutPush struct {
	ProjectGuid         *string `json:"projectGuid"`
	ProjectName         *string `json:"projectName"`
	ProjectNo           *string `json:"projectNo"`
	Reason              *string `json:"reason"`
	LogoutType          *int64  `json:"logoutType"`
	WinBidderName       *string `json:"winBidderName"`
	WinBidderCreditCode *string `json:"winBidderCreditCode"`
}
