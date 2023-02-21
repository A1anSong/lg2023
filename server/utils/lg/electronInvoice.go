package lg

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/go-resty/resty/v2"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func get_sign(senid string, nonce string, timestamp string, content string) string {
	param := "a=services&l=v1&p=open&k=" + global.GVA_CONFIG.Insurance.NNAppKey + "&i=" + senid + "&n=" + nonce + "&t=" + timestamp + "&f=" + content
	h := hmac.New(sha1.New, []byte(global.GVA_CONFIG.Insurance.NNAppSecret))
	h.Write([]byte(param))
	hmacCode := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(hmacCode)
}

func NNSendRequest(method string, content string) ([]byte, error) {
	tokenTime := time.UnixMilli(global.GVA_CONFIG.Insurance.NNTokenTime)
	if time.Since(tokenTime).Hours() > 30*time.Hour.Hours() {
		err := nnRequestAccessToken()
		if err != nil {
			return []byte(""), err
		}
	}

	senid := strings.Replace(uuid.NewV4().String(), "-", "", -1)
	nonce := rand.Intn(89999999) + 10000000
	timestamp := time.Now().Unix()
	rand.Seed(timestamp)
	sign := get_sign(senid, strconv.Itoa(nonce), strconv.FormatInt(timestamp, 10), content)

	requestUrl := global.GVA_CONFIG.Insurance.NNRequestUrl + "?senid=" + senid + "&nonce=" + strconv.Itoa(nonce) + "&timestamp=" + strconv.FormatInt(timestamp, 10) + "&appkey=" + global.GVA_CONFIG.Insurance.NNAppKey
	client := resty.New()
	client.Header["Content-Type"] = []string{"application/json"}
	client.Header["X-Nuonuo-Sign"] = []string{sign}
	client.Header["accessToken"] = []string{global.GVA_CONFIG.Insurance.NNAccessToken}
	client.Header["userTax"] = []string{global.GVA_CONFIG.Insurance.NNTaxNo}
	client.Header["method"] = []string{method}
	resp, err := client.R().
		SetBody(content).
		Post(requestUrl)
	if err != nil {
		return []byte(""), err
	}
	if resp.StatusCode() == http.StatusOK {
		return resp.Body(), nil
	} else {
		return []byte(""), errors.New("电子发票系统响应失败")
	}
}

func nnRequestAccessToken() error {
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
