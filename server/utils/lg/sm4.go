package lg

import (
	"encoding/base64"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/tjfoc/gmsm/sm4"
)

func Sm4Encrypt(data []byte) string {
	cbcMsg, _ := sm4.Sm4Cbc([]byte(global.GVA_CONFIG.Insurance.AppSecret), data, true)
	return base64.StdEncoding.EncodeToString(cbcMsg)
}

func Sm4Decrypt(msg []byte) (string, error) {
	cbcDec, err := sm4.Sm4Cbc([]byte(global.GVA_CONFIG.Insurance.AppSecret), msg, false)
	return string(cbcDec), err
}
