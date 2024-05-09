package cmd

// CmderModel 分发命令
type CmderModel struct {
}

var Cmder = CmderModel{}

func (c CmderModel) InitCmd(cmdList []string) {
	cmdTwo := cmdList[1]
	// 获取第二个参数
	switch cmdTwo {
	case "db":
		c.Makemigrations()
	}
}
