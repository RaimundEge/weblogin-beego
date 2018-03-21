package controllers

import (
	"github.com/astaxie/beego"
)

type Home struct {
	beego.Controller
}

func (c *Home) Get() {
	user := c.GetSession("user")
	if user == nil {
		c.Data["user"] = ""		
	} else {
		c.Data["user"] = user
	}
	c.Data["page"] = "home"
	c.Data["message"] = ""
	c.Layout = "layout.html"
	c.TplName = "home.tpl"	
}

type Content struct {
	beego.Controller
}

func (c *Content) Get() {
	user := c.GetSession("user")
	if user == nil {
		c.Data["user"] = ""
		c.Data["page"] = "home"		
		c.TplName = "home.tpl"		
	} else {
		c.Data["user"] = user
		c.Data["page"] = "content"		
		c.TplName = "content.tpl"
	}
	c.Data["message"] = ""
	c.Layout = "layout.html"
}