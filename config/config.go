package config

type Server struct {
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Loggers Loggers `mapstructure:"logger" json:"logger" yaml:"logger"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
}
