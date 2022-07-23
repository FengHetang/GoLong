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
// go 语言诞生的比较晚，  web开发逐渐主流  高并发
//多线程- 每个线程占用的内存比较多 并且系统切换开销很大  上千  绿程/轻量级宪曾  -  用户态的线程

// python 有两种编程模式   1.多线程和多线程来进行并行编程 2. 使用携程进程并发编程，很多的库都是基于多线程和多进程开发

func print_() {
	fmt.Println("bobby")
}

func main() {
	go print_()
	// 主死从随
	time.Sleep(time.Second * 2)
	fmt.Println("imooc")
}
