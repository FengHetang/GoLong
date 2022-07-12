/**
 * @Author: 戈亓
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/7/12 9:28
 */

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()

	// 配置路由
	r.GET("/ping", func(context *gin.Context) {
		context.String(200, "值：%v", "你好gin")
	})

	r.GET("/news", func(context *gin.Context) {
		context.String(http.StatusOK, "我是新闻")
	})

	r.POST("/add", func(context *gin.Context) {
		context.String(200, "post请求，主要用于添加数据")
	})

	r.DELETE("/del", func(context *gin.Context) {
		context.String(200, "Delete请求，主要用于删除数据")
	})

	r.PUT("/put", func(context *gin.Context) {
		context.String(200, "这是put请求，用编辑数据")
	})
	r.Run() //启动一个路由 默认再8080 端口  r.run("8000")  更改端口
}
