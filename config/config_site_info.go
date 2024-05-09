package config

type SiteInfo struct {
	CreatedAt   string `mapstructure:"created_at" json:"created_at" yaml:"created_at"`       // 建站日期
	BeiAn       string `mapstructure:"bei_an" json:"bei_an" yaml:"bei_an"`                   // 备案
	Title       string `mapstructure:"title" json:"title" yaml:"title"`                      // 网站标题
	QQImage     string `mapstructure:"qq_image" json:"qq_image" yaml:"qq_image"`             // QQ图片
	Version     string `mapstructure:"version" json:"version" yaml:"version"`                // 版本
	Email       string `mapstructure:"email" json:"email" yaml:"email"`                      // 邮箱
	WechatImage string `mapstructure:"wechat_image" json:"wechat_image" yaml:"wechat_image"` // 微信图片
	Name        string `mapstructure:"name" json:"name" yaml:"name"`                         // 名字
	Job         string `mapstructure:"job" json:"job" yaml:"job"`                            // 工作
	Addr        string `mapstructure:"addr" json:"addr" yaml:"addr"`                         // 地址
	Slogan      string `mapstructure:"slogan" json:"slogan" yaml:"slogan"`                   // slogan
	SloganEn    string `mapstructure:"slogan_en" json:"slogan_en" yaml:"slogan_en"`          // 英文slogan
	Web         string `mapstructure:"web" json:"web" yaml:"web"`                            // web站点
	BiliBiliUrl string `mapstructure:"bilibili_url" json:"bilibili_url" yaml:"bilibili_url"` // web站点
	GiteeUrl    string `mapstructure:"gitee_url" json:"gitee_url" yaml:"gitee_url"`          // web站点
	GithubUrl   string `mapstructure:"github_url" json:"github_url" yaml:"github_url"`       // web站点
}
