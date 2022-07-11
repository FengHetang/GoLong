/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/11 22:09
 */

// go 语言不支持继承，但是有办法能达到同样的效果 组合

package main

import "fmt"

type Course struct {
	Teacher Teacher // 方法一
	//  Teacher 方法二  若出现重名的字段 先找 Course 中的字段 	 所以需要指明字段    Course.Teacher
	Name  string
	Price int
	Url   string
}

func (c Course) courseInfo() {
	fmt.Printf("课程名：%s,价格：%d，讲师信息：%s,%d,%s", c.Name, c.Price, c.Teacher.Name, c.Teacher.Age, c.Teacher.Title)
}

type Teacher struct {
	Name  string
	Age   int
	Title string
}

func (t Teacher) teacherInfo() {
	fmt.Printf("姓名：%s,年龄：%d，职称：%s,", t.Name, t.Age, t.Title)
}

func main() {
	// go 语言的继承   组合
	t := Teacher{
		Name:  "bobby",
		Age:   18,
		Title: "教授",
	}
	c := Course{
		Name:    "Django",
		Teacher: t,
		Price:   100,
		Url:     "http:www.imooc.com",
	}
	c.courseInfo()

}
