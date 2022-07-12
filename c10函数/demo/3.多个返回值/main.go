/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/13 2:29
 */

package main

import (
	"fmt"
	"strconv"
)

func MuyltiResult(numberOnt int, numberTwo int) (int, string) {
	sum := numberOnt + numberTwo
	return sum, strconv.Itoa(sum)
}
func main() {
	fmt.Println(MuyltiResult(100, 200))
}

//go 是强类型的编程语言，输入参数后，返回值都要需要指定数据类型
// 具有多个返回值时需要带 （） ，比如（int,string）
// 函数调用时参数的舒徐要和定义参数时保持一致，有几个输入参数，就要传入几个参数；有几个返回值，就要接收几个返回值
