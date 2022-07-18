/**
 * @Author: 戈亓
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2022/7/19 3:24
 */

package main

import "fmt"

func main() {
	//什么是切片
	// 数组有一个很大的问题： 大小确定，不能修改，go语言提供了切片 slice 相当于动态数组
	courses := []string{"django", "scrapy", "tornado"}
	fmt.Printf("%T\n", courses) // []string

	//第二种： make
	course := make([]string, 5) //make函数 第一个参数是 类型，第二个参数是长度
	fmt.Printf("%T\n", course)  // []string
	fmt.Println(len(course))    // 5
	// slice 对标的就是python 中的 list

	//第三种方法： 通过数组变成一个切片
	var cours = [5]string{"django", "scrapy", "tornado", "python", "asyncio"}
	subcours := cours[1:4]       // python中的用法叫做切片 go语言中切片是一种数据结构
	fmt.Println(subcours)        // [scrapy tornado python]
	fmt.Printf("%T\n", subcours) // []string

	//第四种方法： new
	subcours2 := *new([]int)
	fmt.Printf("%T\n", subcours2)

	// 切片传递是引用传递
	// slice 很重要 很常用
}
