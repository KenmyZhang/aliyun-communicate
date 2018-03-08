package main

import (
	"encoding/json"
	"fmt"

	"github.com/KenmyZhang/aliyun-communicate"
)

var (
	gatewayUrl      = "http://dysmsapi.aliyuncs.com/"
	accessKeyId     = "LTAIbTnPbawglLIQ"
	accessKeySecret = ""
	phoneNumbers    = "13544285**2"
	signName        = "坤Kenmy"
	templateCode    = "SMS_82045083"
	templateParam   = "{\"code\":\"1234\"}"
)

func main() {
	smsClient := aliyunsmsclient.New(gatewayUrl)
	result, err := smsClient.Execute(accessKeyId, accessKeySecret, phoneNumbers, signName, templateCode, templateParam)
	fmt.Println("Got raw response from server:", string(result.RawResponse))
	if err != nil {
		panic("Failed to send Message: " + err.Error())
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	if result.IsSuccessful() {
		fmt.Println("A SMS is sent successfully:", resultJson)
	} else {
		fmt.Println("Failed to send a SMS:", resultJson)
	}
}
