package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_admin "go-mage-admin/app/admin"
	"go-mage-admin/app/admin/model"
	"go-mage-admin/app/core/route"
	"go-mage-admin/app/mage"
	"net/http"
	"strconv"
	"strings"
)

// init bootstrap初始化时调用 自动注册路由
func init() {
	route.Register(&Grid{}, _admin.Name)
}

// Grid grid相关功能
type Grid struct {
	// 继承控制器基类
}

// ConfigResetPOST 重置自定义列
func (g *Grid) ConfigResetPOST(c *gin.Context) {
	config := model.ConfigGrid{}
	session := sessions.Default(mage.AppGinContext)
	uid, _ := strconv.Atoi(fmt.Sprintf("%v", session.Get("uid")))
	code := mage.AppGinContext.DefaultPostForm("code", "")
	if code != "" {
		flag := config.Reset(uid, code)
		if flag {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "ok",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "重置配置失败！",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数code必填！",
		})
	}
}

// ConfigSavePOST 保存自定义列
func (g *Grid) ConfigSavePOST(c *gin.Context) {
	config := model.ConfigGrid{}
	session := sessions.Default(mage.AppGinContext)
	uid, _ := strconv.Atoi(fmt.Sprintf("%v", session.Get("uid")))
	code := mage.AppGinContext.DefaultPostForm("code", "")
	configData, _ := mage.AppGinContext.GetPostFormArray("config[]")
	if code != "" {
		flag := config.Update(uid, code, strings.Join(configData, ","))
		if flag {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "ok",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "保存配置失败！",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "参数code必填！",
		})
	}

}
