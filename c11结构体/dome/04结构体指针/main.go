/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/9 12:58
 */

package main

import "fmt"

type People struct {
	Name string
	Age  int
}

func main() {
	peo := new(People)
	fmt.Println(peo == nil) // false  结构体中的属性不是指针，可以直接调用

	peo.Name = "张三"
	fmt.Println(peo) // &{张三 0}
	peo1 := peo
	peo1.Name = "李四"
	fmt.Println(peo1, peo) // &{李四 0} &{李四 0}

	//peo2 := *People
	//peo2 = &People{"smallming", 17}
	//fmt.Println(peo2)

}
