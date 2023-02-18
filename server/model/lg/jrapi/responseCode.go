package jrapi

type ResponseCode int

const (
	FAILED              ResponseCode = -1
	SUCCESS             ResponseCode = 0
	MissingParam        ResponseCode = 1001
	NotAuthorized       ResponseCode = 1002
	SignCheckFailed     ResponseCode = 1003
	DecryptFailed       ResponseCode = 1004
	MissingServiceParam ResponseCode = 1005
)

func (r ResponseCode) String() string {
	switch r {
	case FAILED:
		return "系统异常"
	case SUCCESS:
		return "success"
	case MissingParam:
		return "缺少系统必填参数"
	case NotAuthorized:
		return "未授权"
	case SignCheckFailed:
		return "签名校验失败"
	case DecryptFailed:
		return "解密失败"
	case MissingServiceParam:
		return "缺少业务必填参数"
	default:
		return ""
	}
}
