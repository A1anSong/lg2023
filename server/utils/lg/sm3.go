package lg

import (
	"crypto/hmac"
	"encoding/base64"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/tjfoc/gmsm/sm3"
)

func SM3Sign(data string) string {
	h := hmac.New(sm3.New, []byte(global.GVA_CONFIG.Insurance.AppSecret))
	h.Write([]byte(data))
	sum := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(sum)
}
