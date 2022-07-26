/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/22 1:40
 */

package main

import (
	"fmt"
	"time"
)

//goroutine

// python、Java、c++ 多线程和多线程编程
// go 语言诞生的比较晚，web2.0  开发逐渐主流  高并发
//多线程的问题：每个线程占用的内存比较多 并且系统切换开销很大  每个电脑的线程是有限的   绿程/轻量级线程  -主流说法叫协程 -> 用户态的线程
// go语言从一开始就没有打算让我们去实例化一个线程 - 协程
// nodejs python3.6 也开始支持协程
// python 有两种编程模式   1.多线程和多线程来进行并行编程 2. 使用携程进程并发编程，很多的库都是基于多线程和多进程开发
// python有大量的库 但是在使用携程的时候 有的库不支持使用

func print_() {
	//fmt.Println("bobby")
	for {
		fmt.Println("1") //
	}
}

func main() {
	// go的协程和python的协程对比
	// 1.开启一百万个协程
	// 2.使用的简单性

	for i := 0; i < 10000; i++ {
		go func(n int) {
			fmt.Println(n)
			time.Sleep(time.Second)
		}(i)
	}

	// 开启100个线程
	//for i := 0; i < 100; i++ {
	//	// 闭包 特性
	//	go func(n int) {
	//		for {
	//			fmt.Println(n)
	//			time.Sleep(time.Second) // time.sleep 的目的是为了 把打印的值看到
	//		}
	//	}(i) // 将i 传入  生成不同线程的时候 已经将i拷贝进入参数栈，所以各个不同的线程打印的值不同
	//}

	//方法二使用   匿名函数
	//go func() {
	//	for {
	//		fmt.Println("22222")
	//	}
	//}()

	// 方法一：
	// 关键字 go   + 函数 表示该函数进入协程运行
	// go print_() // 主协程中的开辟出来的子协程

	// 主死从随
	time.Sleep(time.Second * 30) // 主线程  只有当主线程开启的时候 子线程才能运行 如果主线程关闭 则子线程也会关闭
	fmt.Println("imooc")
}
