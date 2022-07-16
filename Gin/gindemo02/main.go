/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/12 11:47
 */

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title  string `json:"title"` // 结构体字段  标签的用法
	Desc   string
	Contet string
}

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "值：%v", "首页")
	})

	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "hello",
		})
	})

	//
	r.GET("/join", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"msg":     "你好，gin--22",
		})
	})
	// 返回结构体
	r.GET("/jion3", func(c *gin.Context) {
		a := Article{
			Title:  "我是标题",
			Desc:   "描述",
			Contet: "测试内容",
		}
		c.JSON(200, a)
	})
	r.GET("/jsonp", func(c *gin.Context) {
		a := Article{
			Title:  "我是标题",
			Desc:   "描述",
			Contet: "测试内容",
		}
		c.JSONP(200, a)
	})

	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"success": true,
			"msg":     "你好  gin",
		})
	})
	r.Run()
}
