package admin

import (
	"github.com/gin-contrib/multitemplate"
	"go-mage-admin/app/core/controller"
)

type Admin struct {
}

const name = "admin"

func (admin *Admin) LoadLayout() multitemplate.Renderer {
	c := &controller.Abstract{}
	r := c.LoadLayout(name, true)
	return r
}
