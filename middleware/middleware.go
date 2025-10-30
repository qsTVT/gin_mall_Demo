package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Cors 处理跨域请求的中间件，用于解决前端与后端跨域调用API的问题
func Cors() gin.HandlerFunc {
	// 返回gin框架的HandlerFunc类型，符合中间件接口要求
	return func(c *gin.Context) {
		// 1. 获取当前请求的HTTP方法（如GET、POST、OPTIONS等）
		method := c.Request.Method
		// 2. 获取请求头中的Origin字段，即前端请求的来源域名
		origin := c.Request.Header.Get("Origin")

		// 3. 收集请求头中所有的键（如Authorization、Content-Type等），用于后续处理
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}

		// 4. 将收集到的请求头键拼接成字符串，格式化为跨域相关的头信息字符串
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			// 若有请求头，则拼接为"access-control-allow-origin, access-control-allow-headers: 头1,头2..."
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers: %s", headerStr)
		} else {
			// 若无请求头，则使用默认字符串
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}

		// 5. 若请求有来源（origin不为空），则设置跨域相关的响应头
		if origin != "" {
			// 设置允许跨域的来源为所有域名（*表示通配符，允许任意域）
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			// 重复设置允许跨域的来源（与上一行作用相同，可能为冗余代码）
			c.Header("Access-Control-Allow-Origin", "*")
			// 允许的HTTP请求方法（覆盖电商常用操作：查询、提交、修改、删除等）
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
			// 允许前端请求中携带的头信息（包含认证、数据类型等必要头）
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept,Origin,Host,Connection,Accept-Encoding,Accept-Language,DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Pragma")
			// 允许前端获取的响应头（暴露额外的头信息给前端）
			c.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			// 预检请求（OPTIONS）的结果缓存时间（172800秒=2天），减少重复预检
			c.Header("Access-Control-Max-Age", "172800")
			// 是否允许跨域请求携带凭证（如cookie、token），此处设置为不允许
			c.Header("Access-Control-Allow-Credentials", "false")
			// 设置响应数据格式为JSON
			c.Set("content-type", "application/json")
		}

		// 6. 若请求方法为OPTIONS（浏览器跨域前的预检请求），直接返回成功响应
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		// 7. 继续执行后续的中间件或路由处理函数
		c.Next()
	}

}
