/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/3 15:54
 */

package main

import "fmt"

func main() {
	num := 11
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	score := 85
	if score >= 90 {
		fmt.Println("优")
	} else if score >= 80 {
		fmt.Println("良")
	} else if score >= 60 {
		fmt.Println("一般")
	} else {
		fmt.Println("不及格")
	}

	// nums 仅限于该if 语句中
	if nums := 13; nums%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println(nums)
	}

}
