package http_proxy_middleware

import (
	"github.com/gin-gonic/gin"
)

// HTTPAccessModelMiddleware 匹配接入方式 基于请求信息
//请求信息与服务列表做匹配关系,匹配到苏需要的服务,拿出配置放到上下文里,可以根据匹配到的服务做负载均衡反向代理权限校验等
func HTTPAccessModelMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
