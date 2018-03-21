package routers

import (
	"DemoBeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.Home{})
    beego.Router("/home", &controllers.Home{})
    beego.Router("/content", &controllers.Content{})
    beego.Router("/login", &controllers.Login{})
    beego.Router("/logout", &controllers.Logout{})
    beego.Router("/register", &controllers.Register{})
}
