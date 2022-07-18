/**
 * @Author: 戈亓
 * @Description:
 * @File:  mian
 * @Version: 1.0.0
 * @Date: 2022/7/18 21:59
 */

package main

import "fmt"

func swap(a *int, b *int) {
	// 用于交换  a 和 b
	c := *a
	*a = *b
	*b = c
}

func main() {
	// 什么是指针
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println(a, b) // 交换不成功
	// 指针 - 对于内存来说，每一个字节其实都是有地址 - 通过16进制打印出来
	fmt.Printf("%p\n", &a) // 0x11c140a8
	// 现在有一种特殊的变量类型，这个变量只能保存地址值
	var ip *int // 这个变量就只能保存地址类型这种值
	ip = &a

	//  如果要修改指针指向的变量的值 用法也比较特殊
	*ip = 30
	fmt.Println(a)
	//如何定义指针变量，如果修改指针变量指向的内存种的值。
	//通过指针去取值的时候不知道应该去取多大的连续内存空间
	fmt.Printf("ip所指向的内存空间地址是：%p，内存中的值是：%d", ip, *ip)
	// 数组是值传递 数组中有100万个 对于这种一般采用切片来传递
	// 在python中 list和dict这种传递 都是引用类型

	// 指针还可以指向数组，指向数组的指针  数组是值类型

	//arr := [3]int{1, 2, 3}
	//var ips *[3]int = &arr
	// 指针数组   数组里面有很多指针
	//var ptrs [3]*int // 创建能够存放三个指针变量的数组
	//// 指针 很多时候 都是函数参数的时候指明的类型
	//// 指针的默认值 nil
	//if ptrs != nil {
	//
	//}
	// python 和java 屏蔽指针

	//make 函数 new nil
}
