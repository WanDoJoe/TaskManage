package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)
/*一个默认的控制器*/
type MainController struct {
	beego.Controller
}
type ShortResult struct {
	UrlShort string
	UrlLong  string
}
var (
	urlcache cache.Cache
)

func init() {
	urlcache, _ = cache.NewCache("memory", `{"interval":0}`)
}
func (this *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	//this.Data["json"]=map[string]interface{}{"result":"success","message":"登录成功"}
	//this.ServeJSON()
	////return
	//this.Ctx.WriteString("{\"result\":\"success\",\"message\":\"登录成功\"}")

	//this.Ctx.Output.Body([]byte("shorturl"))
	//var result ShortResult
	//shorturl := this.Input().Get("shorturl")
	//result.UrlShort = shorturl
	//if urlcache.IsExist(shorturl) {
	//	result.UrlLong = urlcache.Get(shorturl).(string)
	//} else {
	//	result.UrlLong = ""
	//}
	//this.Data["json"] = result
	//this.ServeJSON()
}
