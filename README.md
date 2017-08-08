# 短信服务 Go语言开发 
采用的是阿里云-云通信-短信服务，原来的阿里大于品牌已经升级为“阿里云 · 云通信”。在阿里官方文档上只提供了java、php、python、nodejs的样例程序，没有提供go语言版本的，所以这里就补充一下用Go实现的样例程序，仅供参考

# 用途
用户注册、找回密码、用户身份验证、验证码登录等等

# 使用说明
将其中的accessKeyId、accessKeySecret、phoneNumbers、signName、templateCode、templateParam替换成你的就可以直接使用了
	package main 

	import (
		"fmt"
		"github.com/KenmyZhang/aliyun-communicate/app"
	)
	var (
		gatewayUrl = "http://dysmsapi.aliyuncs.com/"		
		accessKeyId = "*******"
		accessKeySecret = "*************"
		phoneNumbers = "1354*********"
		signName = "KenmyZhang"
		templateCode = "SMS_82045083"
		templateParam = "{\"code\":\"1234\"}"
	)

	 func main() {
	 	smsClient := app.NewSmsClient(gatewayUrl)
	 	if result, err := smsClient.Execute(accessKeyId, accessKeySecret, phoneNumbers, signName, templateCode, templateParam); err != nil {
	 		fmt.Println("error:%v", err.Error())
	 	} else {
	 		for key, value := range result {
	 			 fmt.Println("key:%v", key, " value:",value)
	 		}
	 	}
		
	}
## 参考http详解链接
https://help.aliyun.com/document_detail/56189.html?spm=5176.doc55288.6.567.O7dDSP

## 若有疑问，请联系QQ：1027837952
