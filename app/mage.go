package app

import (
	"fmt"
	"go-mage-admin/app/core"
	"time"
)

type Mage struct {
	//版本号
	version string
	//是否已安装
	isInstalled  bool
	isProduction bool
	//App
	app core.App
	//缓存管理
	cache core.Cache
	//全局变量数组
	Registry map[string]interface{}
	//events 实现观察者
	Events map[string]interface{}
	//缓存对象
}

// GetVersion 获取版本号
func (mage *Mage) GetVersion() string {
	return mage.version
}

// Log 日志函数统一入口
func (mage *Mage) Log() {
	return
}

// LogException 日常日志
func (mage *Mage) LogException(msg string) {
	return
}

// PrintException 打印异常
func (mage *Mage) PrintException() {
	return
}

// SetIsDeveloperMode 开发者模式设置[报错和日志等]
func (mage *Mage) SetIsDeveloperMode(flag bool) {
	core.AppConfig.Site.DeveloperMode = flag
	return
}

// SetIsProductionMode 生产模式设置
func (mage *Mage) SetIsProductionMode(flag bool) {
	mage.isProduction = flag
	return
}

// Run 程序入口
func (mage *Mage) Run(options map[string]interface{}) {
	profiler := &core.Profiler{}
	profiler.Start("mage::run")
	//设置时区
	cstZone := time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone

	app := &core.App{}
	app.Run(options)
	profiler.Stop("mage::run")
	fmt.Println("Run hello world!")
	return
}
