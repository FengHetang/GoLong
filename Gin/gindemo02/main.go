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
)

type Article struct {
	Title  string
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
	r.GET("/join", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"msg":     "你好，gin--22",
		})
	})

	r.GET("/jion3", func(c *gin.Context) {
		a := Article{
			Title:  "我是标题",
			Desc:   "描述",
			Contet: "测试内容",
		}
		c.JSON(200, a)
	})
	r.Run()
}
