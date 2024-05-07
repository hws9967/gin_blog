package config

import "fmt"

type System struct {
	Env  string `mapstructure:"env" json:"env" yaml:"env"`    // 环境值 区分开发环境测试环境
	IP   string `mapstructure:"ip" json:"ip" yaml:"ip"`       // 项目运行的ip
	Port int    `mapstructure:"port" json:"port" yaml:"port"` // 项目运行的端口
}

func (s System) GetAddr() string {
	return fmt.Sprintf("%s:%d", s.IP, s.Port)
}
