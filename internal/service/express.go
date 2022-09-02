package service

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"shippo-server/utils"
	"strconv"
	"strings"
)

type ExpressService struct {
	*Service
}

func NewExpressService(s *Service) *ExpressService {
	return &ExpressService{s}
}

func (t *ExpressService) GenerateMsgDigest(msgData string, timestamp int64, checkWord string) string {
	str := url.QueryEscape(msgData + strconv.FormatInt(timestamp, 10) + checkWord)
	fmt.Printf("generateMsgDigest->str:%+v\n", str)

	h := md5.New()
	h.Write([]byte(str))
	msgDigest := base64.StdEncoding.EncodeToString(h.Sum(nil))
	fmt.Printf("generateMsgDigest->msgDigest:%+v\n", msgDigest)

	return msgDigest
}

func (t *ExpressService) SfSearchRoutes(msgData string, msgDigest string, timestamp int64) {

	var data = url.Values{}
	data.Add("partnerID", "")
	data.Add("requestID", utils.GenerateToken())
	data.Add("serviceCode", "EXP_RECE_SEARCH_ROUTES")
	data.Add("timestamp", strconv.FormatInt(timestamp, 10))
	data.Add("msgDigest", msgDigest)
	data.Add("msgData", url.QueryEscape(msgData))

	dataStr, _ := url.QueryUnescape(data.Encode())
	fmt.Printf("ExpressService->SfSearchRoutes->dataStr:%+v\n", dataStr)

	request, _ := http.NewRequest("POST", "https://bspgw.sf-express.com/std/service", strings.NewReader(dataStr))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)

	var result = new(struct {
		ApiErrorMsg   string `json:"apiErrorMsg"`
		ApiResponseID string `json:"apiResponseID"`
		ApiResultCode string `json:"apiResultCode"`
		ApiResultData string `json:"apiResultData"`
	})

	json.Unmarshal(bytes, &result)

	fmt.Printf("ExpressService->SfSearchRoutes->Unmarshal:%+v\n", result)

}
