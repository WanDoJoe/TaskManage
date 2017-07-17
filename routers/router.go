package routers

import (
	"beegoDemo/controllers"
	"github.com/astaxie/beego"
)
/*
请求的json方式
jsonPatams={
"params":{},
"sysinifo":{}
}
*/
func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login",&controllers.LoginController{},"*:Login")
	beego.Router("/addDemo",&controllers.LoginController{},"*:AddDemo")
	beego.Router("/findbyid",&controllers.LoginController{},"*:FindById")
	beego.Router("/findbylogin",&controllers.LoginController{},"*:FindByLogin")//FindByLogin
	beego.Router("/findlistuser",&controllers.LoginController{},"*:FindByListUser")//FindByListUser
}
