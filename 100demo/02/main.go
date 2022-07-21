/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/20 22:15
 */

// 题目：企业发放的奖金根据利润提成。利润(I)低于或等于10万元时，奖金可提10%；利润高于10万元，低于20万元时，低于10万元的部分按10%提成，
//高于10万元的部分，可提成7.5%；20万到40万之间时，高于20万元的部分，可提成5%；40万到60万之间时高于40万元的部分，可提成3%；
//60万到100万之间时，高于60万元的部分，可提成1.5%，高于100万元时，超过100万元的部分按1%提成，从键盘输入当月利润I，求应发放奖金总数？

package main

import "fmt"

func main() {

	var i float32
	fmt.Println("请输入企业净利润：(万元)")
	fmt.Scanln(&i)
	//
	arr := []float32{100, 60, 40, 20, 10, 0}
	rat := []float32{0.01, 0.015, 0.03, 0.05, 0.075, 0.1}

	var sum float32 = 0

	for index := 0; index <= 5; index++ {
		if i > arr[index] {
			sum += (i - arr[index]) * rat[index]
			i = arr[index] //  将输入值赋值为arr 数组中的下一个 然后在次循环 计算下一部分利息
		}
	}

	fmt.Println("总利润是：", sum, "万元")
}
