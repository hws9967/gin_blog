package config

type Config struct {
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"`
}

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     string `yaml:"show_line"`
	LogInConsole string `yaml:"log_in_console"`
}

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}
