/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/9 12:32
 */

package main

import "fmt"

//方法（methods）是与对象实例绑定的特殊函数，是面向对象程序设计中的基本概念，用来执行对象的某些操作，比如获取值，改变值等其他操作，所以结构体能用来实现面向对象程序中的“类” （class）的概念

type MyInt struct {
	Number int
}

// 为MyInt 定义了三个方法  SayHello SetNumber SayNumber
// 方法的调用也使用 “.” 的形式
// 若要修改结构体的属性，则需要传入指针， 比如 func(m *MyInt)

func (m MyInt) SayHello() {
	fmt.Println("Hello ,this funcation named SatHello.")
}
func (m *MyInt) SetNumber(other int) {
	m.Number = other
}

func (m MyInt) SayNumber() {
	fmt.Println(m.Number)
}

func main() {
	var my MyInt
	my.Number = 1
	my.SayHello()  // Hello ,this funcation named SatHello.
	my.SayNumber() // 1
	my.SetNumber(5)
	my.SayNumber() // 5
}
