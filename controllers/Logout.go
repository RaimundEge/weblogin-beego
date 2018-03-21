package controllers

import (
	"github.com/astaxie/beego"
)

type Logout struct {
	beego.Controller
}

func (c *Logout) Get() {
	c.SetSession("user", nil)
	c.Data["user"] = ""
	c.Data["message"] = "You have been logget out"
	c.Data["page"] = "home"
	c.Layout = "layout.html"
	c.TplName = "home.tpl"	
}