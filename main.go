package main

import (
	"gin_blog/core"
	"gin_blog/global"
)

func main() {
	global.CONFIG = core.ConfInit()
	global.LOG = core.InitLogger()
	global.DB = core.Gorm()
}
