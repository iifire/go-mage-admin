package admin

import (
	"github.com/gin-contrib/multitemplate"
	"go-mage-admin/app/core/controller"
)

type Admin struct {
}

const Name = "admin"

func (admin *Admin) LoadLayout(filename string, extraTpl []string) multitemplate.Renderer {
	c := &controller.Abstract{}
	r := c.LoadLayout(Name, true, filename, extraTpl)
	return r
}
