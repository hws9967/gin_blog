package config

type Config struct {
	Mysql  Mysql   `yaml:"mysql"`
	Logger Loggers `yaml:"logger"`
	System System  `yaml:"system"`
}
type Server struct {
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Loggers Loggers `mapstructure:"logger" json:"logger" yaml:"logger"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
}
