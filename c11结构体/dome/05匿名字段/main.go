/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/9 15:08
 */

//结构体中 有数据类，没有变量名称的的字段，被称为匿名字段

package main

import "fmt"

type Student struct {
	Name string
	University
}

type University struct {
	Name     string
	Location string
}

func main() {
	var std = new(Student)
	std.Name = "张三"
	//匿名字段具有和主结构体相同的字段Name 初始化时要用多级 “.” 来引用
	std.University.Name = "清华大学"
	std.Location = "北京"
	fmt.Println(std)
}
