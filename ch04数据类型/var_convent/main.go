/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/6/28 12:18
 */

package main

import "fmt"

func main() {
	//1、 基本数据类型转换
	a := int(3.0)
	fmt.Println(a)
	// 在go语言中不支持变量间的隐式转换
	//1、 变量间类型转换不支持
	//var b int = 5.0 // 常量到变量会进行隐式转换
	c := 5.1
	fmt.Printf("%T\n", c)
	var d int = int(c)
	fmt.Println(d)

}
