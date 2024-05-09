package config

type QiNiu struct {
	AccessKey string  `mapstructure:"access-key" json:"access-key" yaml:"access-key"`
	SecretKey string  `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`
	Bucket    string  `mapstructure:"bucket" json:"bucket" yaml:"bucket"` //存储桶
	CDN       string  `mapstructure:"cdn" json:"cdn" yaml:"cdn"`          //访问图片地址前缀
	Zone      string  `mapstructure:"zone" json:"zone" yaml:"zone"`       //存储地区
	Size      float64 `mapstructure:"size" json:"size" yaml:"size"`       //存储大小限制，单位MB
}
