/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/13 2:23
 */

package main

import "fmt"

func NameResult(numberOne, numberTwo int) (result int) {
	result = numberTwo + numberOne
	return
}

// 队友具有相同类型的多个联系输入参数，可以仅一次类型声明
// 可以给返回值命名
// 命名返回值要作为函数的返回值返回给调用者时，在retrun 语句后不用接变量

func main() {
	sumResult := NameResult(12, 20)
	fmt.Println(sumResult)
}
