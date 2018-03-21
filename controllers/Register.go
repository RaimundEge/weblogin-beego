package controllers

import (
	"github.com/astaxie/beego"
	"DemoBeego/models"
)

type Register struct {
	beego.Controller
}

func (c *Register) Get() {
	if (c.GetSession("user") == nil) {
		c.Data["user"] = ""
		c.Data["message"] = "You must be logged in first"
		c.Data["page"] = "home"
		c.TplName = "home.tpl"	
	} else {
		c.Data["user"] = c.GetSession("user").(string)
		c.Data["message"] = ""
		c.Data["page"] = "register"
		c.TplName = "register.tpl"	
	}
	c.Layout = "layout.html"
}

func (c *Register) Post() {
	c.Data["user"] = c.GetSession("user")
	name := c.GetString("username")
	beego.Info("seeking: " + name)
	if models.FindUser(name) != "" {
		c.Data["message"] = "Username not available"
		c.Data["page"] = "register"
		c.TplName = "register.tpl"
	} else {		
		fullname := c.GetString("fullname")
		password := c.GetString("password")
		models.AddUser(name, fullname, password)
		c.Data["message"] = "New user registered: " + fullname
		c.Data["page"] = "content"
		c.TplName = "content.tpl"
	}
	c.Layout = "layout.html"
}
