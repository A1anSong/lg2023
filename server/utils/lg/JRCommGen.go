package lg

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi/jrrequest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lg/jrapi/jrresponse"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"time"
)

func GenJRRequest(formatData interface{}) (req jrrequest.JRRequest, err error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	requestId := uuid.NewV4().String()
	resJsonData, err := json.Marshal(formatData)
	if err != nil {
		return
	}
	data := Sm4Encrypt(resJsonData)
	req = jrrequest.JRRequest{
		AppKey:    global.GVA_CONFIG.Insurance.AppKey,
		Timestamp: timestamp,
		RequestId: requestId,
		Signature: SM3Sign("appKey=" + global.GVA_CONFIG.Insurance.AppKey + "&data=" + data + "&requestId=" + requestId + "&timestamp=" + timestamp),
		Version:   "1.0.0",
		Data:      data,
	}
	return
}

func GenJRResponse(formatData interface{}) (res jrresponse.JRResponse, err error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	requestId := uuid.NewV4().String()
	resJsonData, err := json.Marshal(formatData)
	if err != nil {
		return
	}
	data := Sm4Encrypt(resJsonData)
	res = jrresponse.JRResponse{
		Code:      int(jrapi.SUCCESS),
		Msg:       jrapi.SUCCESS.String(),
		AppKey:    global.GVA_CONFIG.Insurance.AppKey,
		Timestamp: timestamp,
		RequestId: requestId,
		Signature: SM3Sign("appKey=" + global.GVA_CONFIG.Insurance.AppKey + "&data=" + data + "&requestId=" + requestId + "&timestamp=" + timestamp),
		Data:      data,
	}
	return
}
