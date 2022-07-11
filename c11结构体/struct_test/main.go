/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/10 1:18
 */

package main

import (
	"GoLong/c11结构体/new_pkg"
	"fmt"
	"unsafe"
)

type Course struct {
	Name  string
	Price int
	Url   string
}

// 函数的接收者  把
func (c Course) printCourseInfo() {
	fmt.Printf("课程名:%s,课程的价格：%d,课程的地址：%s\n", c.Name, c.Price, c.Url)
}

func (c *Course) setPrice(price int) {
	c.Price = price
}

// 结构体的方法只能和结构体再一个包里面

func main() {
	//go语言不支持面向对象
	//面向对象的三个基本特征 		1、封装  2、继承  3、多态  4、方法重载 	5、抽象基类  go语言支持前三条
	// 定义struct go语言没有class 这个概念，所以对于很多人来说会少很多面向对象抽象的概念

	//实例化
	var c new_pkg.Course = new_pkg.Course{
		Name:  "django",
		Price: 5,
		Url:   "http://www.imooc.com",
	}

	fmt.Println(c.Name, c.Price, c.Url) //django 5 http://www.imooc.com

	// 大小写在go语言中的重要性 可见性
	// 一个包中的变量或者结构体 如果首字母是小写，那么对于另一个包 不可见
	//结构体定义的 名称 以及属性首字母大写很重要

	//2、第二种方式   顺序形式
	c2 := new_pkg.Course{"scrapy", 111, "http:www.imiic.com"}
	fmt.Println(c2.Name, c2.Price, c2.Url) // scrapy 111 http:www.imiic.com

	//3、 如果一个指向结构体的指针,
	//

	c3 := &new_pkg.Course{"TORNADO", 100, "http://www.imoc.com"}
	fmt.Printf("%T\n", c3) // 输出 	*new_pkg.CourseT  指针类型
	// *c3.Name  会指向new_pkg.Course中Name字段，但是该字段 不是指针类型，即 可以使用一下方法
	fmt.Println((*c3).Name)                //指针指向 c3   输出 TORNADO
	fmt.Println(c3.Name, c3.Price, c3.Url) // 输出TORNADO 100 http://www.imoc.com
	//go语言的语法糖，go语言内部会将c3.Name 转换成(*c3).Name
	// go语言在借鉴动态语言的特性
	//另一个根本的原因 - go语言的指针是受限的

	//4、零值  如果不给结构体赋值，go语言会默认给每个字段猜采用默认值
	c4 := new_pkg.Course{}
	fmt.Println(c4.Price) // 输出 0

	// 5、多种零值初始化结构体
	var c5 new_pkg.Course = new_pkg.Course{}
	var c6 new_pkg.Course
	var c7 *new_pkg.Course = new(new_pkg.Course) // 指针的实例化方式  new（）
	var c9 *new_pkg.Course = &new_pkg.Course{}
	//var c8 *new_pkg.Course

	fmt.Println("零值初始化")
	fmt.Println(c5.Price) // 0
	fmt.Println(c6.Price) // 0
	fmt.Println(c7.Price) // 0
	//fmt.Println(c8.Price)  // 异常指针如果只申明 不复制，默认值是nil c6不是指针是结构体的类型
	// slice map 默认值都是nil
	//可以像c9 一样声明
	fmt.Println(c9.Price) // 0

	//6、结构体是值类型
	c10 := new_pkg.Course{"scrapy", 100, "http://www.imooc.com"}
	c11 := c10
	fmt.Println(c10) //	{scrapy 100 http://www.imooc.com}
	fmt.Println(c11) //	{scrapy 100 http://www.imooc.com}
	c10.Price = 200
	fmt.Println(c10) // {scrapy 200 http://www.imooc.com}  修改c10的属性，并没有影响到c11
	fmt.Println(c11) // {scrapy 100 http://www.imooc.com}

	//7、结构体的大小 占用内存的大小，可以使用 sizeof 来查看 对象占用的类型
	fmt.Println(unsafe.Sizeof(c10)) // 20

	//// 8、slice的大小
	//type slice struct {
	//	array unsafe.Pointer //底层数组的地址
	//	len   int            // 长度
	//	cap   int            // 容量
	//}
	//
	//s1 := []string{"django", "tornado", "scrapy", "celery"}
	//fmt.Println(unsafe.Sizeof(s1)) // 12

	//达到了封装数据和封装方法的效果
	c12 := Course{"scrapy", 110, "http:www.imooc.com"}
	Course.printCourseInfo(c12) // 效果一致
	//c12.printCourseInfo()
	c12.setPrice(120)
	c12.printCourseInfo()
	
}
