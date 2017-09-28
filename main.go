package main

import (
	"fmt"
	"github.com/15515731931/aliyun-communicate/config"
	"github.com/KenmyZhang/aliyun-communicate/app"
)

func main() {
	//read config
	config := config.GetConfig("config.json")
	//create SMSClient
	smsClient := app.NewSmsClient(config)
	//execute
	if result, err := smsClient.Execute(); err != nil {
		fmt.Println("error:", err.Error())
	} else {
		for key, value := range result {
			fmt.Println("key:", key, " value:", value)
		}
	}

}
