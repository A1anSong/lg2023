package jrresponse

type JRResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	AppKey    string `json:"appKey,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	RequestId string `json:"requestId,omitempty"`
	Signature string `json:"signature,omitempty"`
	Data      string `json:"data"`
}
