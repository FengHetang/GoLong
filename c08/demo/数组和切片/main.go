/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/14 17:32
 */

package main

import (
	"fmt"
	"reflect"
)

//	匿名函数
var opList = func(number [4]int) {
	// 第二个值和类型
	fmt.Println(number[1], reflect.TypeOf(number[1]))
	// 长度
	fmt.Println(len(number))
	// 切片和类型
	fmt.Println(number[1:], reflect.TypeOf(number[1:]))
	//数组的遍历
	for index, one := range number {
		fmt.Println(index, one)
	}
	for i := 0; i < len(number); i++ {
		fmt.Println(i, number[i])
	}
}

var opSlice = func(name []string) []string {
	// 第二个值和类型
	fmt.Println(name[1], reflect.TypeOf(name[1]))
	//切片的遍历
	for index, one := range name {
		fmt.Println(index, one)
	}
	//
	name = append(name, "xiewei")
	return name
}

func main() {
	var number [4]int = [...]int{1, 2, 3, 4}
	opList(number)
	var name []string = []string{"Go", "Python", "Java", "C++", "C#"}
	fmt.Println(opSlice(name))
}

// 数组的长度是固定的，切片可以自动扩充容量（长度），append() 方法
//切片可以走动扩充容量，到那时底层并不是动态数组或者数组指针，而是通过指针指向底层数组，所以经常食用数组来创建切片
//切片是引用类型，所以对切片初始化的时候可以采用显示的方法对切片赋值，也可以使用make 关键字
