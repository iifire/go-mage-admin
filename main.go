package main

import (
	"go-mage-admin/app"
)

func main() {
	mage := &app.Mage{}
	options := map[string]interface{}{}
	//TODO... 运行参数设置
	mage.Run(options)
}
