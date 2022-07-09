/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/9 12:45
 */

package main

import "fmt"

//组合：结构体嵌套结构体，组合即可实现类的继承

type ViewName struct {
	Name string
	ViewOther
}

func (v ViewName) SayName() {
	fmt.Println(v.Name)
}

type ViewOther struct {
	Value string
}

func (v ViewOther) SayValue() {
	fmt.Println(v.Value)
}

func main() {
	//初始化结构体 ViewName
	var viewName ViewName
	viewName.Name = "张三"
	// 调用嵌套在 ViewName 结构体中的 另一个结构体 ViewOther 的属性
	viewName.Value = "value"
	viewName.SayName()
	viewName.SayValue()
}
