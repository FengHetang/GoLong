/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/12 22:31
 */

package main

import "fmt"

// numberOne he numberTwo 是函数的两个参数，数据类型是int型
// 该函数有一个返回值，数据类型是int型
func SumNumber(numberOne int, numberTwo int) int {
	return numberOne + numberTwo
}

func main() {
	SumAdd := SumNumber(1, 2)
	fmt.Println(SumAdd)
}
