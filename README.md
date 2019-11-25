package main

import (
	"fmt"

	"github.com/yujinxiang/xssfilter/filter"
)

type Student struct {
	Name  string
	Grade string
}

func main() {

	//结构体
	stu := Student{Name: "<张三>", Grade: "<女>"}
	err := filter.XssFilter(&stu)
	fmt.Println(fmt.Sprintf("%+v", stu), err)

	//map
	mdata := map[string]interface{}{"key1": "", "key2": map[string]interface{}{"key2_1": "<value2_1>"}}
	err = filter.XssFilter(mdata)
	fmt.Println(fmt.Sprintf("%+v", stu), err)

	//slice
	mdata := []string{"", ""}
	err = filter.XssFilter(mdata)
	fmt.Println(fmt.Sprintf("%+v", stu), err)

}
