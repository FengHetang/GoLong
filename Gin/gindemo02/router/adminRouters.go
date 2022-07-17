/**
 * @Author: 戈亓
 * @Description:
 * @File:  adminRouters
 * @Version: 1.0.0
 * @Date: 2022/7/17 22:07
 */

package router

import "github.com/gin-gonic/gin"

func AdminRouter(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", func(c *gin.Context) {
			c.String(200, "后台首页")
		})
		adminRouters.GET("/user", func(c *gin.Context) {
			c.String(200, "用户列表")
		})
		adminRouters.GET("/user/add", func(c *gin.Context) {

		})
	}
}
