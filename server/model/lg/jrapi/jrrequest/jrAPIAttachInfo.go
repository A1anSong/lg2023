package jrrequest

type JRAPIAttachInfo struct {
	AttachType *string `json:"attachType,omitempty"`
	AttachName *string `json:"attachName,omitempty"`
	AttachUrl  *string `json:"attachUrl,omitempty"`
}
