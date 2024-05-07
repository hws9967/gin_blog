package main

import (
	"fmt"
	"gin_blog/core"
	"gin_blog/global"
)

func main() {
	core.ConfInit()
	fmt.Println(global.CONFIG)
}
