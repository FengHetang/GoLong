/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/20 21:55
 */

// 有四个数字：1、2、3、4，能组成多少个互不相同且无重复数字的三位数？各是多少？

package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			for k := 1; k <= 4; k++ {
				if (i != k) && (i != j) && (j != k) {
					fmt.Println(i, j, k)
				}
			}
		}
	}
}
