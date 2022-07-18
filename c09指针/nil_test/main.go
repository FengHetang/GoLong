/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/18 22:31
 */

package main

import "fmt"

func main() {
	// make he new 函数
	// new 函数用法
	//var p *int // 声明了一个变量p 但是变量没有初始值 没有内存
	//*p = 10
	// 默认值 int byte rune float bool string 这些类型都有默认值
	// 指针 切片 map 接口 这些的默认值都是nil 理解为none
	var a = 10
	fmt.Println(a)

	// 对于指针或者说其他的默认值是0 的情况来说如何一开始申明的时候就分配内存
	var p *int = new(int)
	*p = 10
	//var
}
