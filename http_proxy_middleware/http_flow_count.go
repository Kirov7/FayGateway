package http_proxy_middleware

import (
	"fmt"
	"github.com/Kirov7/FayGateway/dao"
	"github.com/Kirov7/FayGateway/middleware"
	"github.com/Kirov7/FayGateway/public"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"time"
)

func HTTPFlowCountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		service, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}
		serviceDetail := service.(*dao.ServiceDetail)

		// 1. 全站统计
		totalCounter, err := public.FlowCounterHandler.GetCounter(public.FlowTotal)
		if err != nil {
			middleware.ResponseError(c, 4001, err)
			c.Abort()
			return
		}
		totalCounter.Increase()

		dayCount, _ := totalCounter.GetDayData(time.Now())
		fmt.Printf("totalCounter qps:%v,dayCount:%v", totalCounter.QPS, dayCount)

		// 2. 服务统计
		serviceCounter, err := public.FlowCounterHandler.GetCounter(public.FlowCountServicePrefix + serviceDetail.Info.ServiceName)
		if err != nil {
			middleware.ResponseError(c, 4002, err)
			c.Abort()
			return
		}
		serviceCounter.Increase()

		dayServiceCount, _ := serviceCounter.GetDayData(time.Now())
		fmt.Printf("serviceCounter qps:%v,dayCount:%v", serviceCounter.QPS, dayServiceCount)

		// 3. 租户统计
		//appCounter, err := public.FlowCounterHandler.GetCounter(public.FlowCountAppPrefix)
		//if err != nil {
		//	middleware.ResponseError(c, 4003, err)
		//	c.Abort()
		//	return
		//}
		//appCounter.Increase()

		c.Next()
	}
}
