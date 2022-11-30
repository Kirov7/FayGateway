package http_proxy_middleware

import (
	"github.com/Kirov7/FayGateway/dao"
	"github.com/Kirov7/FayGateway/middleware"
	"github.com/Kirov7/FayGateway/reverse_proxy"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// HTTPReverseProxyMiddleWare 匹配接入方式 基于请求信息
func HTTPReverseProxyMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		serviceIf, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}
		serviceDetail := serviceIf.(*dao.ServiceDetail)
		lb, err := dao.LoadBalancerHandler.GetLoadBalancer(serviceDetail)
		if err != nil {
			middleware.ResponseError(c, 2002, err)
			c.Abort()
			return
		}
		trans, err := dao.TransporterHandler.GetTrans(serviceDetail)
		if err != nil {
			middleware.ResponseError(c, 2003, err)
			c.Abort()
			return
		}
		proxy := reverse_proxy.NewLoadBalanceReverseProxy(c, lb, trans)
		proxy.ServeHTTP(c.Writer, c.Request)
		//创建ReverseProxy
		//使用ReverseProxy.ServeHTTP(c.request, c.Respone)服务
	}
}
