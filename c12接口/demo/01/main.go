/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/13 21:00
 */

package main

import "fmt"

// 接口是对其他类型行为的抽象和概括，是一组方法的集合。接口制定了对象的行为，是Go语言中的非常重要的数据类型

type Controller interface {
	SayHello()
	SayNumber(int)
	SayHi()
}

type DefaultController struct {
}

func (d DefaultController) SayHello() {
	fmt.Println("Hello  world!")
}
func (d DefaultController) SayNumber(number int) {
	fmt.Println(fmt.Sprintf("%d", number))
}

func (d DefaultController) SayHi() {
	fmt.Println("Say Hi")
}

func main() {
	var d DefaultController
	var c Controller
	c = d
	c.SayHello()
	c.SayHi()
	c.SayNumber(123)
}

//定义一个接口Controller，并在其中定义了三个接口 SayHello SayNumber SayHi
// 定义一个结构体DefaultController,并具体实现了这三个方法
// DefaulController 实现了接口 Controller
