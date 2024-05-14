package config

type QiNiu struct {
	Enable    bool    `mapstructure:"enable" json:"enable" yaml:"enable"` // 是否启用七牛云存储
	AccessKey string  `mapstructure:"access_key" json:"access_key" yaml:"access_key"`
	SecretKey string  `mapstructure:"secret_key" json:"secret_key" yaml:"secret_key"`
	Bucket    string  `mapstructure:"bucket" json:"bucket" yaml:"bucket"`  //存储桶
	CDN       string  `mapstructure:"cdn" json:"cdn" yaml:"cdn"`           //访问图片地址前缀
	Zone      string  `mapstructure:"zone" json:"zone" yaml:"zone"`        //存储地区
	Prefix    string  `mapstructure:"prefix" ,json:"prefix" yaml:"prefix"` // 前缀
	Size      float64 `mapstructure:"size" json:"size" yaml:"size"`        //存储大小限制，单位MB
}
