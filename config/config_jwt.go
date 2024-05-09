package config

type Jwy struct {
	Secret  string `mapstructure:"secret" json:"secret" yaml:"secret"`    // 密钥
	Expires int    `mapstructure:"expires" json:"expires" yaml:"expires"` // 过期时间
	Issuer  string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`    //
}
