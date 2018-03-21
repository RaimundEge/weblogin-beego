package controllers

import (
	"github.com/astaxie/beego"
	"DemoBeego/models"
)

type Login struct {
	beego.Controller
}

func (c *Login) Get() {
	if (c.GetSession("user") == nil) {
		c.Data["user"] = ""
	} else {
		c.Data["user"] = c.GetSession("user").(string)
	}
	c.Data["message"] = ""
	c.Data["page"] = "login"
	c.TplName = "login.tpl"
	c.Layout = "layout.html"
}

func (c *Login) Post() {
	name := c.GetString("username")
	// beego.Info("seeking: " + name)
	password := c.GetString("password")
	
	fullname := models.VerifyUser(name, password)
	if fullname == "" {
		c.Data["user"] = ""
		c.Data["message"] = "Username/password incorrect"
		c.Data["page"] = "login"
		c.TplName = "login.tpl"
	} else {		
		c.SetSession("user", fullname)
		c.Data["user"] = c.GetSession("user").(string)
		c.Data["message"] = fullname + ": Welcome back!"
		c.Data["page"] = "content"
		c.TplName = "content.tpl"
	}
	c.Layout = "layout.html"
}