package config

type Logger struct {
	Level        string `mapstructure:"level" json:"level" yaml:"level"`                            // 级别
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日志前缀
	Director     string `mapstructure:"director" json:"director"  yaml:"director"`                  // 日志文件夹
	ShowLine     bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`                // 显示行
	LogInConsole bool   `mapstructure:"log_in_console" json:"log_in_console" yaml:"log_in_console"` // 输出控制台
}
