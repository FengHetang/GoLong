/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/13 21:23
 */

package main

import "fmt"

// 接口是一个协议   其实就是一组方法的集合

type Programmer interface {
	Coding() string //方法只是声明
	Debug() string
}

// java里面一种类型只要继承一个接口才行，如果你继承了这个几口的话，那么这个接口里面的所有方法你必须要全部实现
type Pythoner struct {
}

func (p Pythoner) Coding() string {
	fmt.Println("python 开发者")
	return "Python 开发者"
}

// 对于pythoner这个结构体来说 实现任何方法第，但是如果不全部实现Coding Debug 的话 那 Pythoner 就不是一个 Progrommer类型
// 1.pythoner 本身自己就是一个类型，那我何必在意自己是不是Programmer
// 2、  封装、继承、多态
// 3、interface：在go语言中接口是一种类型 是一种抽象类型
func main() {

}
