package api

import (
	"github.com/astaxie/beego"
)

type InitController struct {
	beego.Controller
}

func (c *InitController) Get() {
	c.Data["json"] = string("{tip:'success',data:'name'}")
	c.ServeJSON()
}
