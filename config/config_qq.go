package config

import "fmt"

type QQ struct {
	AppID    string `mapstructure:"app-id" json:"app-id" yaml:"app-id"`
	Key      string `mapstructure:"key" json:"key" yaml:"key"`
	Redirect string `mapstructure:"redirect" json:"redirect" yaml:"redirect"` //登入之后的回调地址
}

func (q QQ) GetPath() string {
	if q.Key == "" || q.AppID == "" || q.Redirect == "" {
		return ""
	}
	return fmt.Sprintf("https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=%s&redirect_uri=%s", q.AppID, q.Redirect)
}
