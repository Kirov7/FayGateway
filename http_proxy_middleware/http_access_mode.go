package http_proxy_middleware

import (
	"fmt"
	"github.com/e421083458/FayGateway/dao"
	"github.com/e421083458/FayGateway/middleware"
	"github.com/e421083458/FayGateway/public"
	"github.com/gin-gonic/gin"
)

// HTTPAccessModeMiddleware 匹配接入方式 基于请求信息
//请求信息与服务列表做匹配关系,匹配到苏需要的服务,拿出配置放到上下文里,可以根据匹配到的服务做负载均衡反向代理权限校验等
func HTTPAccessModeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		service, err := dao.ServiceMangerHandler.HTTPAccessMode(c)
		if err != nil {
			middleware.ResponseError(c, 1001, err)
			c.Abort()
			return
		}
		fmt.Println("matched service", public.Obj2Json(service))
		c.Set("service", service)
		c.Next()
	}
}
