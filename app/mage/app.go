// Package mage 全局变量定义
package mage

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AppDb 数据库配置
var AppDb = make(map[string]*gorm.DB)

// AppGin Gin.Engine全局变量
var AppGin *gin.Engine

// AppGinContext AppGin Gin.Context全局变量
var AppGinContext *gin.Context

// AppModules 当前启用模块列表
var AppModules = make(map[string]interface{})
