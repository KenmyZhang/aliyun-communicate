// @see `./sms-lib/sms-lib.go`.
// The client of the sms service of Aliyun(阿理云短信服务客户端) which can be named as aliyunsmsclient or simply smsclient, or smsserviceclient.
package aliyunsmsclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/KenmyZhang/aliyun-communicate/sms-lib"
)

type SmsClient struct {
	Request    *aliyunsmslib.ALiYunCommunicationRequest
	GatewayUrl string
	Client     *http.Client
}

func New(gatewayUrl string) *SmsClient {
	smsClient := new(SmsClient)
	smsClient.Request = &aliyunsmslib.ALiYunCommunicationRequest{}
	smsClient.GatewayUrl = gatewayUrl
	smsClient.Client = &http.Client{}
	return smsClient
}

func (smsClient *SmsClient) Execute(accessKeyId, accessKeySecret, phoneNumbers, signName, templateCode, templateParam string) (result map[string]interface{}, err error) {
	var endpoint string
	if err = smsClient.Request.SetParamsValue(accessKeyId, phoneNumbers, signName, templateCode, templateParam); err != nil {
		return
	}
	if endpoint, err = smsClient.Request.BuildSmsRequestEndpoint(accessKeySecret, smsClient.GatewayUrl); err != nil {
		return
	}

	request, _ := http.NewRequest("GET", endpoint, nil)
	response, err := smsClient.Client.Do(request)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	defer response.Body.Close()

	err = json.Unmarshal(body, &result)

	return
}
