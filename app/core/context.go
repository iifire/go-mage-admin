package core

import "github.com/gin-gonic/gin"

// Context 继承gin.Context便于重构
type Context struct {
	gin.Context
}
