package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

const TABLENAME_DEMO  ="demo"
type Demo struct {
	Id int
	Name string
	Address string
	Password string
}

func (d *Demo) TableName() string{
 return TableName(TABLENAME_DEMO)
}
//添加
func DemoAdd(demo *Demo)(int64, error){
	fmt.Println("DemoAdd.name="+demo.Name)
	if demo.Name==""{
		return 0,fmt.Errorf("用户名不能为空")
	}

	return orm.NewOrm().Insert(demo)
}
//按照id寻找
func DemoFindById(id int)(*Demo, error){
	task := &Demo{
		Id: id,
	}

	err := orm.NewOrm().Read(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}
func FindByNameAndPassword(name string ,password string)(*Demo ,error){
	//task := &Demo{
	//	Name:name,
	//	Password:password,
	//}

	demo := new(Demo)

	err := orm.NewOrm().QueryTable(TableName(TABLENAME_DEMO)).Filter("name", name).Filter("password",password).One(demo)
	if err != nil {
		return nil, err
	}
	return demo, nil
}
var pas string
var pa string
func DemoFindList(page, pageSize int) ([]*Demo, int64) {

	pas =fmt.Sprintf("%d",page)//,pageSize
	pa =fmt.Sprintf("%d",pageSize)
	fmt.Println(pas+"--"+pa)

	offset := (page - 1) * pageSize

	list := make([]*Demo, 0)

	query := orm.NewOrm().QueryTable(TableName(TABLENAME_DEMO))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}

//寻找list
//func DemoFindList(page, pageSize int, filters ...interface{}) ([]*Demo, int64) {
//	offset := (page - 1) * pageSize
//
//	tasks := make([]*Demo, 0)
//
//	query := orm.NewOrm().QueryTable(TableName(TABLENAME_DEMO))
//	if len(filters) > 0 {
//		l := len(filters)
//		for k := 0; k < l; k += 2 {
//			query = query.Filter(filters[k].(string), filters[k+1])
//		}
//	}
//	total, _ := query.Count()
//	query.OrderBy("id").Limit(pageSize, offset).All(&tasks)
//
//	return tasks, total
//}
