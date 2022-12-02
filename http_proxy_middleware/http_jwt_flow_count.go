package http_proxy_middleware

import (
	"fmt"
	"github.com/Kirov7/FayGateway/dao"
	"github.com/Kirov7/FayGateway/middleware"
	"github.com/Kirov7/FayGateway/public"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func HTTPJwtFlowCountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//service, ok := c.Get("service")
		//if !ok {
		//	middleware.ResponseError(c, 2001, errors.New("service not found"))
		//	c.Abort()
		//	return
		//}
		//serviceDetail := service.(*dao.ServiceDetail)

		app, ok := c.Get("app")
		if !ok {
			c.Next()
			return
		}
		appInfo := app.(*dao.App)
		// 租户统计
		appCounter, err := public.FlowCounterHandler.GetCounter(public.FlowAppPrefix + appInfo.AppID)
		if err != nil {
			middleware.ResponseError(c, 2001, err)
			c.Abort()
			return
		}
		appCounter.Increase()

		if appInfo.Qpd > 0 && appCounter.TotalCount > appInfo.Qpd {
			middleware.ResponseError(c, 2002, errors.New(fmt.Sprintf("租户日请求量限流 limit: %v currtent: %v", appInfo.Qpd, appCounter.TotalCount)))
			c.Abort()
			return
		}
		fmt.Println("day count", appCounter.TotalCount)
		c.Next()
	}
}
