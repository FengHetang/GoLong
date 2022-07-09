/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/6/28 11:33
 */

package main

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

func main() {
	// 定义布尔类型
	var gendor = false
	fmt.Println(gendor)

	var age int16 = 18
	fmt.Println(unsafe.Sizeof(age))

	//float类型
	var weight float32 = 71.2
	fmt.Println(weight)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	//为什么 64 为 的flaat 最大值远大于 int64
	//float底层存储和int存储不一样
	//float64 和float 32 占用内存都不一样

	height := 72.3
	fmt.Printf("%T\n", height)
	//默认使用 float64
	aged := 18
	fmt.Printf("%T\n", aged)

	// byte类型
	var a byte = 18
	fmt.Println(a)

	b := 'b'
	fmt.Println(reflect.TypeOf(b + 1))
	fmt.Printf("b+1=%c", b+1)

}
