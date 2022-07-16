/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/16 16:14
 */

package main

import (
	"fmt"
	"time"
)

func TimeUsage() {
	new := time.Now()
	// 获取年
	fmt.Println(new.Year())
	// 获取月
	fmt.Println(new.Month())
	// 获取日期
	fmt.Println(new.Date())
	// 获取天
	fmt.Println(new.Day())
	// 获取小时
	fmt.Println(new.Hour())
	// 获取分钟
	fmt.Println(new.Minute())
	// 获取秒
	fmt.Println(new.Second())
	// 获取毫秒
	fmt.Println(new.Unix())
	// 获取纳秒
	fmt.Println(new.UnixNano())
}

// TimeOperate 获取时间的基本属性
func TimeOperate() {
	start := time.Now()
	time.Sleep(1 * time.Second)
	// 两个时间差
	fmt.Println(time.Now().Sub(start))
	// 格式化
	fmt.Println(start.Format("2006-01-02 15:04:05"))
	// 截取
	fmt.Println(start.Round(time.Second))
	fmt.Println(start.Truncate(time.Second))
	stringTime := "1991-12-25 19:00:00"
	birthday, _ := time.ParseInLocation("2006-01-02 15:04:05", stringTime, time.Local)
	fmt.Println(birthday.String())
}

func main() {
	//TimeUsage()
	TimeOperate()
}
