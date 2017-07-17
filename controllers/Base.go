package controllers

import (
	"github.com/widuu/gojson"
	"fmt"
)

var sysinfo string
func GetParams(json string) *gojson.Js{
	params:=gojson.Json(json).Get("params")
	sysinfo=gojson.Json(json).Get("sysinfo").Tostring()
	return params
}

func GetJSON(JSON *gojson.Js,key string) string{
	fmt.Println(JSON)
	params:=JSON.Get(key).Tostring()
	return params
}
