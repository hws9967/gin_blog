package main

import (
	"fmt"
	"gin_blog/core"
	"gin_blog/global"
)

func main() {
	core.InitConf()
	fmt.Println(global.Config.System.Addr())
	fmt.Println(global.Config.SiteInfo)
	fmt.Println(global.Config.QiNiu)
}
