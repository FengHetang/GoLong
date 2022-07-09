/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/9 12:13
 */

package main

import "fmt"

func main() {
	// go语言的关键词  type
	// 1、 给一个类型 定义别名   rune byte
	//type byte = uint8
	// 为什么会有byte 不使用 uint8  为了强调 我么你现在处理的对象是字节类型
	// 这种别名还是为了代码的可读性 本质上 任上后uint8 代码编码阶段 可读性强而已

	type myByte = byte
	var b myByte
	// var b uint8
	fmt.Printf("%T\n", b) // uint8

	// 2、第二种 就是基于一个已有的类型定义一个新的类型
	type myInt int
	var i myInt
	fmt.Printf("%T\n", i) // main.myInt

	// 3、定义结构体
	type Course struct {
		name  string
		price int
	}
	// 4、定义接口
	type Callable interface {
	}
	//5、定义函数别名
	type handle func(str string)
}
