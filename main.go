package main

import (
	_ "beegoDemo/routers"
	"github.com/astaxie/beego"
	"beegoDemo/models"
)
/*主文件*/
/*
"github.com/widuu/gojson"  json解析
gojson.Json(jsonPatams).Get("Objectkey\Arraykey").Get("paramKey").Tostring()

*/

//所有方法名必须大写
func main() {
	models.Init();



	beego.Run()
}

