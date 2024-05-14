package config

type Mysql struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                               // 服务器地址:端口
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                               //:端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                         // 高级配置
	DB           string `mapstructure:"db" json:"db" yaml:"db"`                                     // 数据库名
	User         string `mapstructure:"user" json:"user" yaml:"user"`                               // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 数据库密码
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogLevel     string `mapstructure:"log_level" json:"log_level" yaml:"log_level"`                // 是否开启Gorm全局日志
}

func (m *Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DB + "?" + m.Config
}
