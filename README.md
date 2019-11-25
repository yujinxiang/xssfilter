package main

import (
	"fmt"

	"github.com/yujinxiang/xssfilter"
)

type Student struct {
	Name  string
	Grade string
}

func main() {

	//结构体
	stu := Student{Name: "<张三>", Grade: "<女>"}
	err := xssfilter.XssFilter(&stu)
	fmt.Println(fmt.Sprintf("%+v", stu), err)

	//map
	mdata := map[string]interface{}{"key1": "<value1>", "key2": map[string]interface{}{"key2_1": "<value2_1>"}}
	err = xssfilter.XssFilter(mdata)
	fmt.Println(fmt.Sprintf("%+v", mdata), err)

	//slice
	sdata := []string{"<fasfsf>", "<rrrrrr>"}
	err = xssfilter.XssFilter(sdata)
	fmt.Println(fmt.Sprintf("%+v", sdata), err)

}

output:
{Name:&lt;张三&gt; Grade:&lt;女&gt;} <nil>
map[key1:&lt;value1&gt; key2:map[key2_1:&lt;value2_1&gt;]] <nil>
[&lt;fasfsf&gt; &lt;rrrrrr&gt;] <nil>
