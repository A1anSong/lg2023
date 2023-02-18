package jrrequest

type JRRequest struct {
	AppKey    string `json:"appKey"`
	Timestamp string `json:"timestamp"`
	RequestId string `json:"requestId"`
	Signature string `json:"signature"`
	Version   string `json:"version"`
	Data      string `json:"data"`
}
