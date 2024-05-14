package config

import "fmt"

type System struct {
	Env  string `mapstructure:"env" json:"env" yaml:"env"`    // 环境值 区分开发环境测试环境
	Host string `mapstructure:"host" json:"host" yaml:"host"` // 项目运行的ip
	Port int    `mapstructure:"port" json:"port" yaml:"port"` // 项目运行的端口
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
