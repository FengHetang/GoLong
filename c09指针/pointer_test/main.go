/**
 * @Author: 戈亓
 * @Description:
 * @File:  mian
 * @Version: 1.0.0
 * @Date: 2022/7/18 21:59
 */

package main

import "fmt"

func swap(a int, b int) {
	// 用于交换  a 和 b
	c := a
	a = b
	b = c
}

func main() {
	// 什么是指针
	a := 10
	b := 20
	swap(a, b)
	fmt.Println(a, b)
}
