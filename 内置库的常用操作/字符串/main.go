/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/15 21:00
 */

package main

import (
	"fmt"
	"strings"
)

const Value string = "Go is an open source programmer language that makes it easy to build simple, reliable, and effcient software."

// StingContaine   1.是否包含字符串
func StingContaine(subStrings string) bool {
	return strings.Contains(Value, subStrings)
}

// StringCompare  2.字符串比较
func StringCompare(value string, subString string) int {
	return strings.Compare(value, subString)
}

// StringToUper  4.大小写转换 转为大写
func StringToUper(subStings string) string {
	return strings.ToUpper(subStings)
}

// StringToLower 大小写转换  转为小写
func StringToLower(subStrings string) string {
	return strings.ToLower(subStrings)
}

// StringToTitle
func StringToTitle(substrings string) string {
	return strings.ToTitle(substrings)
}

// ToUpper 当字符串中 不全是字符的时候 参考以下案例
//func ToUpper(s string) string{
//	isASCII, hasLower := true, false
//	for i := 0; i< len(s); i++{
//		c := s[i]
//		if c >= utf8.RuneSelf{
//			isASCII = false
//			break
//		}
//		hasLower = hasLower || (c >= 'a' && c <= 'z')
//	}
//	if isASCII {
//		if !hasLower{
//			return s
//		}
//		b := make( []byte,len(s))
//		for i :=0; i<len(s); i++{
//			c := s[i]
//			if c >= 'a' && c <= 'z' {
//				c -= 'a' - 'A'
//			}
//			b[i] = c
//		}
//		return string(b)
//	}
//	return Map(unicode.ToUpper,s)
//}

// StringsCount  4.统计字串的次数
func StringsCount(subSrings string) int {
	return strings.Count(Value, subSrings)
}

// StringHasPrefix 5.前缀
func StringHasPrefix(substings string) bool {
	return strings.HasPrefix(Value, substings)
}

// StringHasSuffix 后缀
func StringHasSuffix(subStrings string) bool {
	return strings.HasSuffix(Value, subStrings)
}

// StringSplit 6. 字符串分割
func StringSplit(spilt string) []string {
	return strings.Split(Value, spilt)
}

// StringJoin 字符串连接
func StringJoin(subStrings []string) string {
	return strings.Join(subStrings, " ")
}

// StringTrim 7. 索引
func StringTrim(subSrings string) int {
	return strings.Index(Value, subSrings)
}

// StringsTrim 8. 清洗 一般是去除两端的空格
func StringsTrim(values string) string {
	return strings.TrimSpace(values) // 去除空格
}

// StringsReplacer 9.替换操作
func StringsReplacer(valeus string) string {
	newReplacer := strings.NewReplacer("\n", "", "\t", "", " ", "")
	return newReplacer.Replace(valeus)
}

func main() {
	// 1.判断是否包含子字符串
	fmt.Println(StingContaine("Go"))   //true
	fmt.Println(StingContaine("Java")) // false

	// 2. 字符串比较  根据单个字符的ASCII编码来进行的，比如A是编码是 65 B的编码是 66 ，如果A和B 比较 就会返回 -1
	fmt.Println(StringCompare("Java", "Go")) // -1
	fmt.Println(StringCompare("Go", "Java")) // -1
	fmt.Println(StringCompare("A", "B"))     // -1

	// 3. 大小写转换
	fmt.Println(StringToUper("golong, hello world")) // GOLONG, HELLO WORLD
	fmt.Println(StringToLower("Golong"))             // golong
	fmt.Println(StringToTitle("golong,hello world")) //GOLONG,HELLO WORLD

	// 4.统计字串出现的次数
	fmt.Println(StringsCount("Go")) // 1
	fmt.Println(StringsCount("s"))  // 6

	// 5.字符串的前后缀  源码 种是判断字符串截取之后的子字符串是否相等，
	// 源码如下：
	//func HasPrefix(s interface{}, prefix interface{}) {
	//return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
	//}
	//
	//func HasSuffix(s,suffix string) bool {
	//return len(s) >= len(suffix) && s[len(s)- len(suffix):] == suffix
	//}

	fmt.Println(StringHasSuffix("software"))  // false
	fmt.Println(StringHasSuffix("software.")) // true
	fmt.Println(StringHasPrefix("Java"))      // false
	fmt.Println(StringHasPrefix("Go"))        // true

	// 6.分割和连接
	fmt.Println(StringSplit(","), len(StringSplit(",")))      //  [Go is an open source programmer language that makes it easy to build simple  reliable  and effcient software.] 3
	fmt.Println(StringJoin([]string{"Go", "Python", "Java"})) // Go Python Java

	//7. 索引  获取指定字符串 首次出现的位置
	fmt.Println(StringTrim("o")) // 1

	// 8.清洗 去除字符串两端的空格
	fmt.Println(StringsTrim("     hello   word!      "))

	//9.替换操作
	fmt.Println(StringsReplacer(" hello world, \n golong")) //helloworld,golong

}
