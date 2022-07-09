/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/3 16:04结构体指针
 */

package main

import "fmt"

func HelloGolong(langeuage string) {
	if langeuage == "Go" {
		fmt.Println("Hello" + langeuage)
	}
}

func HelloGo(langeuage string) {
	if langeuage == "Go" {
		fmt.Println("Hello" + langeuage)
	} else {
		fmt.Println("Hello ?" + langeuage)
	}
}

//if...else if ...else..
func HelloLanguage(language string) {
	if language == "Go" {
		fmt.Println("Hello Geogle")
	} else if language == "Python" {
		fmt.Println("import this")
	} else {
		fmt.Println("???" + language)
	}
}

// switch ... case
func useswitch(str string) {
	switch str {
	case "A":
		fmt.Println("Score >=", 90)
	case "B":
		fmt.Println("80 <= Score < 90")
	case "C":
		fmt.Println("70 <= Score < 80 ")
	case "D":
		fmt.Println("60 <= Score < 70")
	case "E", "F":
		fmt.Println("50 <= Score < 60 ")
	default:
		fmt.Println("Score < 50")
	}
}

// for 循环

func SayGo(num int) {
	for i := 0; i < num; i++ {
		fmt.Println("Hello Golang")
	}
}

//for 循环 遍历操作
func useForRange(names []string) {
	for index, name := range names {
		fmt.Println(index, name)
	}
}

func main() {
	HelloGolong("Go")
	HelloGo("Go")
	HelloGo("Python")
	HelloLanguage("Go")
	HelloLanguage("Python")
	HelloLanguage("Java")
	useswitch("A")
	useswitch("B")
	useswitch("D")
	useswitch("dad")
	SayGo(3)
	useForRange([]string{"zhao", "zhang", "qian", "sun", "li", ":zhou", "wang"})
}
