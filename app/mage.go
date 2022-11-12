package app

import (
	"fmt"
	"go-mage-admin/app/core"
)

type Mage struct {
	//版本号
	version string
	//是否开发者模式
	isDeveloperMode bool
	//是否已安装
	isInstalled bool
	//App
	app core.App
	//配置
	config core.Config
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

// SetIsDeveloperMode 开发者模式设置
func (mage *Mage) SetIsDeveloperMode(flag bool) bool {
	mage.isDeveloperMode = flag
	return mage.isDeveloperMode
}

// Helper TODO...  获取帮助类利用reflect来实现消耗性能
func (mage *Mage) Helper(name string) interface{} {
	registryKey := "_helper/" + name
	if mage.Registry[registryKey] == nil {
		//抛出异常
		mage.LogException("Helper not exist:" + name)
	}
	return mage.Registry[registryKey]
}

// Run 程序入口
func (mage *Mage) Run(options map[string]interface{}) {
	profiler := &core.Profiler{}
	profiler.Start("mage::run")
	app := &core.App{}
	app.Run(options)
	profiler.Stop("mage::run")
	fmt.Println("Run hello world!")
	return
}
