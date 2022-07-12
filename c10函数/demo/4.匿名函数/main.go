/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/13 2:35
 */

package main

import "fmt"

var anonymousFuncTimes = func(numberOne int) int {
	return numberOne * 10
}

func main() {
	fmt.Println(anonymousFuncTimes(100))
}

// 匿名函数使用与业务逻辑比较简单的场合
// 匿名函数通常赋值给一个变量，使用该变量即可实现匿名函数的调用
// 其他输入参数和返回值的处理方式与一般函数的处理方式一致
