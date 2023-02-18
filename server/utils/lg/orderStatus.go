package lg

import "github.com/flipped-aurora/gin-vue-admin/server/model/lg"

func OrderStatus(order lg.Order) (status string) {
	if order.LogoutID != nil {
		status = "销函"
	} else if order.ClaimID != nil && *order.Claim.AuditStatus == 2 {
		status = "理赔"
	} else if order.RefundID != nil && *order.Refund.AuditStatus == 2 {
		status = "退函"
	} else if order.DelayID != nil && *order.Delay.AuditStatus == 2 {
		status = "延期"
	} else if order.LetterID != nil {
		status = "已开"
	} else {
		status = "未开"
	}
	return
}

func AuditStatus(statusInt int64) (status string) {
	switch statusInt {
	case 1:
		status = "待审"
	case 2:
		status = "通过"
	case 3:
		status = "拒绝"
	default:
		status = ""
	}
	return
}
