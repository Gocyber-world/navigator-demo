package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 测试使用

func PrintHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Print Headers")
		for k, v := range c.Request.Header {
			fmt.Println(k, v)
		}
		c.Next()
	}
}

// 提取scf的header

func GetInfoFromHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("request-id", c.Request.Header.Get("X-Scf-Request-Id"))
		// 访问者ip
		if fip := c.Request.Header.Get("X-Forwarded-For"); fip != "" {
			c.Set("ip", fip)
		} else {
			c.Set("ip", c.ClientIP())
		}
		c.Next()
	}
}
