package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
	"github.com/widuu/gojson"
	"beegoDemo/models"
	"strconv"
)

type JsonParams struct  {
	params Params `json:params`
	sysinfo Sysinfo `json:sysinfo`
}
type Sysinfo struct {
	drevice string  `json:drevice`
}
type Params struct {
	id string `json:id`
	name string `json:name`
}


type LoginController struct {
	beego.Controller
}
var paramsJson_struct JsonParams
var sysinfo_struct Sysinfo
var params_struct Params

//通过子struct的方法来重写post函数，实现自己的逻辑。
func (c *LoginController) Login() {
	//返回响应处理信息
	jsonPatams :=c.GetString("jsonPatams")
	fmt.Println("paramsJson.params="+jsonPatams)
	params_struct.id =gojson.Json(jsonPatams).Get("params").Get("id").Tostring()
	params_struct.name=gojson.Json(jsonPatams).Get("params").Get("name").Tostring()
	fmt.Println("id="+params_struct.id)
	fmt.Println("name="+params_struct.name)
	//发送json
	c.Data["json"] = map[string]interface{}{"success":true,"message":"登录成功"}
	c.ServeJSON(false)
}


var idstr string
//http://127.0.0.1:9999/addDemo?jsonPatams={"sysinfo":{"drevice":"phone"},"params":{"name":"admin1","address":"地址地址地址"}}
func(c *LoginController)AddDemo(){


	jsonPatams :=c.GetString("jsonParams")
	fmt.Println("AddDemo    jsonParams="+jsonPatams)

	name :=gojson.Json(jsonPatams).Get("params").Get("name").Tostring()
	address :=gojson.Json(jsonPatams).Get("params").Get("address").Tostring()
	fmt.Println("name="+name)
	//实例化结构体
	demo  := new(models.Demo)
	//设置参数
	demo.Name=name
	demo.Address=address
	id,_ :=models.DemoAdd(demo)

	if id>0 {
		idstr= fmt.Sprintf("%d", id)
	}else{
		idstr="插入未成功"
	}
	//发送json
	c.Data["json"] = map[string]interface{}{"success":true,"message":idstr}
	c.ServeJSON(false)
}

//http://127.0.0.1:9999/findbyid?jsonParams={"sysinfo":{"drevice":"phone"},"params":{"id":"3"}}
func (c *LoginController)FindById(){
	jsonPatams :=c.GetString("jsonParams")
	fmt.Println("AddDemo    jsonParams="+jsonPatams)

	idstr :=gojson.Json(jsonPatams).Get("params").Get("id").Tostring()
	id, err := strconv.Atoi(idstr)
	if err != nil {
		panic(err)
	}

	demo,error :=models.DemoFindById(id)
	if error != nil {
		panic(error)
	}
	//发送json
	c.Data["json"] = map[string]interface{}{"success":true,"message":"成功","name":demo.Name,"address":demo.Address}
	c.ServeJSON(false)
}
//http://127.0.0.1:9999/findbylogin?jsonParams={"sysinfo":{"drevice":"phone"},"params":{"password":"1234","name":"admin"}}
func (c *LoginController)FindByLogin(){
	jsonPatams :=c.GetString("jsonParams")
	fmt.Println("AddDemo    jsonParams="+jsonPatams)

	name :=gojson.Json(jsonPatams).Get("params").Get("name").Tostring()
	pwd :=gojson.Json(jsonPatams).Get("params").Get("password").Tostring()

	demo,error :=models.FindByNameAndPassword(name,pwd)
	if error !=nil{
		panic(error)
	}
	//发送json
	c.Data["json"] = map[string]interface{}{"success":true,"message":"成功","id":demo.Id,"name":demo.Name,"address":demo.Address}
	c.ServeJSON(false)
}

//http://127.0.0.1:9999/findlistuser?jsonParams={"sysinfo":{"drevice":"phone"},"params":{"page":"1","pageSize":"10"}}
func (this *LoginController) FindByListUser(){

	paramsJson:=this.GetString("jsonParams")
	var js =GetParams(paramsJson)

	//pageStr:=gojson.Json(paramsJson).Get("params").Get("page").Tostring()
	pageStr:=GetJSON(js,"page")
	page,error := strconv.Atoi(pageStr)
	if error != nil{
		fmt.Println("字符串转换成整数失败")
	}
	//pageSizeStr :=gojson.Json(paramsJson).Get("params").Get("pageSize").Tostring()
	var jsps =GetParams(paramsJson)
	pageSizeStr:=GetJSON(jsps,"pageSize")
	pageSize,err := strconv.Atoi(pageSizeStr)
	if pageSize < 0 {
		pageSize = 1
	}
	result, count := models.DemoFindList(page, pageSize)
	fmt.Sprintf("%d",count)
	//jsonStr:=[]string{}
	jsonStr :=[]models.Demo{}
	//for k, v := range result {
	for _, v := range result {
		row :=models.Demo{
			Id:v.Id,
			Name:v.Name,
			Address:v.Address,
			Password:v.Password,
		}
		jsonStr=append(jsonStr,row)
	}

	array,err:=json.Marshal(jsonStr);
	if err!=nil{
		panic(err)
	}
	jsonarray :=string(array)
	fmt.Println(jsonarray)

	//发送json
	this.Data["json"] = map[string]interface{}{"success":true,"message":"成功","list":jsonarray}
	this.ServeJSON(false)
}
