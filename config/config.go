package config

type Server struct {
	Mysql    Mysql    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Local    Local    `mapstructure:"local" json:"local" yaml:"local"`
	Jwy      Jwy      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	QQ       QQ       `mapstructure:"qq" json:"qq" yaml:"qq"`
	Loggers  Loggers  `mapstructure:"logger" json:"logger" yaml:"logger"`
	System   System   `mapstructure:"system" json:"system" yaml:"system"`
	Email    Email    `mapstructure:"email" json:"email" yaml:"email"`
	QiNiu    QiNiu    `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	SiteInfo SiteInfo `mapstructure:"site-info" json:"site-info" yaml:"site-info"`
}
