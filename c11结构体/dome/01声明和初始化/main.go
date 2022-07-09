/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/9 12:18
 */

package main

import "fmt"

//type 结构体名字 struct {
//	字段 字符类型
//}

type Info struct {
	Name       string
	Age        int
	University string // 公有属性   首字母大写
	sex        string // 私有属性  首字母小写
}

//结构体内的字段也被称为这个结构体的属性，属性分为外部可访问（公有，首字母大写，外部库可以访问）和外部不可访问（私有属性，首字母大写，外部库不可以访问）

//结构体的初始化有两种方法 ，如下， 内存空间会分配一个对应的内存地址
func main() {
	// 第一种方法 直接声明
	var info Info
	info = Info{
		Name:       "xiewei",
		Age:        14,
		University: "清华大学",
	}
	// into = {"xiewei",14,"清华大学"}  按照顺序直接赋值即可 上面的方法中可以指定赋值
	//第二种方法  使用new（） 关键字
	var infoTwo = new(Info)
	// 属性或值的访问 使用"." 的形式
	infoTwo.Name = "xiewei"
	infoTwo.Age = 16
	infoTwo.University = "北京大学"
	fmt.Println(info, infoTwo, *infoTwo)
	// 输出结果： {xiewei 14 清华大学 } &{xiewei 16 北京大学 } {xiewei 16 北京大学 }
}
