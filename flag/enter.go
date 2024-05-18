package flag

import (
	sys_flag "flag"
	"github.com/fatih/structs"
)

type Option struct {
	DB   bool
	User string // -u admin  -u user
	ES   bool   // -es create  -es delete
	Dump string
	Load string
}

// Parse 解析命令行参数
func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户")
	es := sys_flag.Bool("es", false, "es操作")
	dump := sys_flag.String("dump", "", "将索引下的数据导出为json文件")
	load := sys_flag.String("load", "", "导入json数据，并创建索引")
	// 解析命令行参数写入注册的flag里
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
		ES:   *es,
		Dump: *dump,
		Load: *load,
	}
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) (f bool) {
	maps := structs.Map(&option)
	for _, v := range maps {
		switch val := v.(type) {
		case string:
			if val != "" {
				f = true
			}
		case bool:
			if val == true {
				f = true
			}
		}
	}
	return f
}

// SwitchOption 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
		return
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
		return
	}
	/*if option.ES {
		global.ESClient = core.EsConnect()
		if option.Dump != "" {
			DumpIndex(option.Dump)
		}
		if option.Load != "" {
			LoadIndex(option.Load)
		}
	}*/
}
