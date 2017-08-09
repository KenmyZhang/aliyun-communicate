package main 

import (
	"fmt"
	"github.com/KenmyZhang/aliyun-communicate/app"
)
var (
	gatewayUrl = "http://dysmsapi.aliyuncs.com/"		
	accessKeyId = "LTAIbTnPbawglLIQ"
	accessKeySecret = ""
	phoneNumbers = "13544285**2"
	signName = "坤Kenmy"
	templateCode = "SMS_82045083"
	templateParam = "{\"code\":\"1234\"}"
)

 func main() {
 	smsClient := app.NewSmsClient(gatewayUrl)
 	if result, err := smsClient.Execute(accessKeyId, accessKeySecret, phoneNumbers, signName, templateCode, templateParam); err != nil {
 		fmt.Println("error:", err.Error())
 	} else {
 		for key, value := range result {
 			 fmt.Println("key:", key, " value:",value)
 		}
 	}
	
}

