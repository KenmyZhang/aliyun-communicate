package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//Config...
type Config struct {
	GatewayUrl      string `json:"gatewayUrl"`
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:accessKeySecret`
	PhoneNumbers    string `json:phoneNumbers`
	SignName        string `json:signName`
	TemplateCode    string `json:templateCode`
	TemplateParam   string `json:-`
}

//read config file
func GetConfig(file string) *Config {

	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read config file <%s> failed,errors:\n%+v", file, err)
	}
	conf := &Config{}
	if err := json.Unmarshal(data, conf); err != nil {
		fmt.Printf("parse config file <%s> failed,errors:\n%+v", file, err)
	}
	return conf
}
