/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/14 17:23
 */

package main

import "fmt"

func prientArray(toPrint [5]string) {
	// []string 是切片 数组作为参数时必须指定大小
	toPrint[0] = "bobby"
	fmt.Println(toPrint) // [bobby scrapy1 tornado1 python+go asyncio]

}

func main() {
	// Go 语言中的数组 和python中的list 可以对应起来理解，
	//slice 和 python 的list更像
	// 静态语言（Go）中的数组： 1、大小确定  2、类型一致
	// 数组的声明

	// 第一种
	//var courses [10] string
	// 最常用的方法
	courses := [5]string{"django", "scrapy", "tornado"} // 初始化

	// 1、修改值  取值  不支持 删除、添加
	fmt.Println(courses[0]) // django
	// 修改值
	courses[0] = "django3"
	fmt.Println(courses) // [django3 scrapy tornado  ]

	// 数组的另一种创建方式
	// 如果初始化的时候没有全部赋值，部分类型数据会显示对应的默认值
	var a = [4]float32{1.0}
	fmt.Println(a) // [1 0 0 0]

	var c = [5]int{'A', 'B'} // [65,66,0,0]  ASCII编码
	fmt.Println(c)

	// 省略号
	d := [...]int{1, 2, 3, 4, 5}
	fmt.Println(d)

	// 初始化 索引为4的值 为100
	e := [5]int{4: 100} //      [0 0 0 0 100]
	fmt.Println(e)

	f := [...]int{0: 1, 4: 1, 9: 100} //  [1 0 0 0 1 0 0 0 0 100]
	fmt.Println(f)

	// 7-2
	// 求数组长度
	fmt.Println(len(f)) // 10

	// 遍历数组
	for i, value := range courses {
		fmt.Println(i, value)
	}

	// for range 求和
	sum := 0
	for _, value := range f {
		sum += value
	}
	fmt.Println(sum) // 102

	// 使用for 语句遍历数组
	sum = 0
	for i := 0; i < len(courses); i++ {
		sum += f[i]
	}
	fmt.Println(sum) // 2

	// 7-3
	// 数组是值类型
	courseA := [3]string{"django", "scrapy", "tornado"}
	courseB := [...]string{"django", "scrapy1", "tornado1", "python+go", "asyncio"}
	// courseA 和CourseB 应该是同一种类型，都是数组类型
	// go语言中 couseA和course都是数组但是不是同一种类型  长度不同、类型不同 均不是同一种类型
	fmt.Printf("%T\n", courseA) // [3]string
	fmt.Printf("%T\n", courseB) // [5]string

	prientArray(courseB) // 传入courseA 报错 因为courseA 长度是3
	fmt.Println(courseB) // [django scrapy1 tornado1 python+go asyncio]  但是prientArray() 中已经修改数据 所以说明 数组作为参数是，实际调用的时候是值传递
	///python 实现引用传递

}
