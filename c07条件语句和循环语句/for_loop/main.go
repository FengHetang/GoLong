/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/3 17:26
 */

package main

import "fmt"

func main() {
	//for  init;condition;post { }   计算 1-10 的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//死循环
	//for {
	//	fmt.Println("111111")
	//}

	//类似于while 用法
	i := 0
	for i < 10 {
		//for ;i < 10 ;  {
		fmt.Println("baby")
		i++
	}

}
