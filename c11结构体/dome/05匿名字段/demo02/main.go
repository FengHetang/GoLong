/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/9 15:16
 */

package main

import "fmt"

type Student struct {
	Name string
	University
}

func (s Student) PrintName() {
	fmt.Println(s.Name)
}

type University struct {
	Name     string
	Location string
}

func (u University) PrintName() {
	fmt.Println(u.Name)
}

func main() {
	var std = new(Student)
	std.Name = "张三"
	std.University.Name = "清华"
	std.Location = "北京"
	fmt.Println(std)

	std.PrintName()
	std.University.PrintName()
}
