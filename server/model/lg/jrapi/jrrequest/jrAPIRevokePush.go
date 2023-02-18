package jrrequest

type JRAPIRevokePush struct {
	OrderNo      *string `json:"orderNo"`
	RevokeOrigin *string `json:"revokeOrigin"`
	RevokeReason *string `json:"revokeReason"`
}
