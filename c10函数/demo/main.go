/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/12 22:24
 */

package main

import "fmt"

func PrintHello() {
	fmt.Println("Hello world!")
}

// 匿名函数
var PrintName = func() {
	fmt.Println("Hello Name")
}

func main() {
	PrintName()
	PrintHello()
}

// 以关键字func 开头定义一个函数，可以有传入参数，也可以没有传入参数，可以有返回值，也可以没有返回值
// 匿名函数，不以func 开头也可以定义一个函数，称为匿名函数。赋予一个变量，通过该变量，即可调用匿名函数，匿名函数多用于简单的处理逻辑
// 函数调用需要带上（）
// 返回值使用 return
// 函数名称大写，表示可导出，小写则表示私有，不可导出
