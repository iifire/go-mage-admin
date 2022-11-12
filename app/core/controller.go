package core

import (
	"github.com/gin-gonic/gin"
	r "go-mage-admin/app/core/route"
)

// Controller 路由解析
type Controller struct {
	ginServer *gin.Engine
}

// Dispatch 路由分发
func (c *Controller) Dispatch() {
	r.InitRouter(c.ginServer)
	return
}
