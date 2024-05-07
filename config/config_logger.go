package config

type Loggers struct {
	Level        string `mapstructure:"level" json:"level" yaml:"level"`                            // 级别
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日志前缀
	Director     string `mapstructure:"director" json:"director"  yaml:"director"`                  // 日志文件夹
	ShowLine     bool   `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // 显示行
	LogInConsole bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // 输出控制台
}
