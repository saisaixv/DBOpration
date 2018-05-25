package routers

import (
	"DBOpration/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/userinfo", &controllers.GetUserInfoController{})
	beego.Router("/keepalive", &controllers.KeepaliveController{})

}
