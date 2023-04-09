package response

type ElogValidateMessage struct {
	Status     *string `json:"status" form:"status"`
	Message    *string `json:"message" form:"message"`
	SubMessage *string `json:"subMessage" form:"subMessage"`
}
