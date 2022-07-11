/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/11 22:23
 */

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//type Info struct {
//	Name   string `json:"name"`          // name是映射成mysql中的varchar类型还是text类型
//	Age    int    `json:"age,omitempty"` // omitenmpy/ 在序列化的时候直接忽略0值 或者空值
//	Gender string	`json:"-"`
//}

type Info struct {
	Name   string `orm:"name,max_length,min_length=5"`
	Age    int    `orm:"age,min=10, max=70"`
	Gender string `orm:"gendor,required"` // 用逗号，隔开 符合规范即可
}

//反射包

func main() {
	//结构体标签
	/*
		结构体的字段除了名字和类型外，还可以有一个可选的标签（tag：）
		它是一个附属于字段的字符串，可以是文档或者其他的重要标记。
		比如我们再解析json或生成json文件时，常用到encoding/json包，
		它提供一些默认标签，例如 omitempty标签可以直接再序列化的时候忽略0值或者空值。
		而-标签的作用是不进行序列化，其效果和直接将结构体中的字段写成小写的效果一样
	*/
	info := Info{
		Name:   "bobby",
		Gender: "男",
	}
	re, _ := json.Marshal(info)
	fmt.Println(string(re)) // 若Info 的age字段  `json:"age,omitempty"`没有声明 则 输出 age 初始值为0

	// 通过反射包去识别  tag
	t := reflect.TypeOf(info)
	fmt.Println("Type", t.Name()) //  Info
	fmt.Println("Kind", t.Kind()) // struct

	// NumField 属性的个数
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i) //获取结构体的每一个字段
		tag := field.Tag.Get("orm")
		fmt.Printf("%d,%v,(%v),tag:'%v\n", i+1, field.Name, field.Type.Name(), tag)
	}
	// 具体的应用绝大部分情况之下是不需要使用到反射的  实际开发的项目中会使用的到
}
