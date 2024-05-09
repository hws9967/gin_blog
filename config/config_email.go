package config

type Email struct {
	Host             string `mapstructure:"host" json:"host" yaml:"host"`
	Port             int    `mapstructure:"port" json:"port" yaml:"port"`
	User             string `mapstructure:"user" json:"user" yaml:"user"`
	Password         string `mapstructure:"password" json:"password" yaml:"password"`
	DefaultFromEmail string `mapstructure:"default-from-email" json:"default-from-email" yaml:"default-from-email"`
	UseSSL           bool   `mapstructure:"use-ssl" json:"use-ssl" yaml:"use-ssl"`
	UserTls          bool   `mapstructure:"user-tls" json:"user-tls" yaml:"user-tls"`
}
