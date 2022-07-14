/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/14 20:48
 */

package main

import "fmt"

//字典是映射类型的一种，是一种搜索速度非常快的数据结构，日常使用频率非常高，是一种无序键值对的集合

var onMap = func(name map[string]int) {
	// 遍历键和值
	for key, value := range name {
		fmt.Println(key, value)
	}
	// 赋值
	name["life"] = 100
	// 判断是否存在key： Go
	if value, ok := name["GO"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("no exeits Go")
	}
	//删除 key:  java
	delete(name, "java")
}

func main() {
	nameMap := make(map[string]int)
	nameMap["java"] = 200
	nameMap["php"] = 100
	nameMap["python"] = 180
	nameMap["js"] = 220
	onMap(nameMap)
}

// map 是引用类型 使用make 初始化
// 无序：输出键的顺序和定义的顺序不一致。
// 可以遍历键和值
// 可以判断是否存在某个键
// 可以删除某个键
// map是引用类型，作为参数传入，在操作过程中会改变其值，请慎重使用
