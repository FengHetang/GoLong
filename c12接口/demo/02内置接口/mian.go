/**
 * @Author: 戈亓
 * @Description:
 * @File:  mian
 * @Version: 1.0.0
 * @Date: 2022/7/13 21:12
 */

package main

import "fmt"

type error interface {
	Error() string
}

type ErrorCode struct {
	Code    int
	Message string
}

func (e ErrorCode) Error() string {
	return fmt.Sprintf("Code :%d, Message: %s", e.Code, e.Message)
}

//func SayError() error {
//	var e ErrorCode
//	e.Code = 400
//	e.Message = "http status code"
//}

func main() {
	//SayError()
}
