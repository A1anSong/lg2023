package lg

import (
	"encoding/json"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func NNRequestAccessToken() error {
	type Token struct {
		AccessToken *string `json:"access_token"`
		ExpiresIn   *int    `json:"expires_in"`
	}
	var token Token
	client := resty.New()
	client.Header["Content-Type"] = []string{"application/x-www-form-urlencoded;charset=UTF-8"}
	resp, err := client.R().
		SetFormData(map[string]string{
			"client_id":     "SD63236305",
			"client_secret": "SDDED2523BED4643",
			"grant_type":    "client_credentials",
		}).
		Post("https://open.nuonuo.com/accessToken")
	if err != nil {
		return err
	}
	if resp.StatusCode() == http.StatusOK {
		err = json.Unmarshal(resp.Body(), &token)
		if err != nil {
			return err
		}
		if token.AccessToken == nil {
			err := errors.New("获取AccessToken失败")
			global.GVA_LOG.Error("获取AccessToken失败", zap.Error(err))
			return err
		} else {
			global.GVA_CONFIG.Insurance.NNAccessToken = *token.AccessToken
			global.GVA_CONFIG.Insurance.NNTokenTime = time.Now().UnixMilli()
			cs := utils.StructToMap(global.GVA_CONFIG)
			for k, v := range cs {
				global.GVA_VP.Set(k, v)
			}
			err = global.GVA_VP.WriteConfig()
			if err != nil {
				return err
			}
			return err
		}
	} else {
		return errors.New("电子发票系统响应失败")
	}
}
