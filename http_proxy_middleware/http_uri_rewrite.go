package http_proxy_middleware

import (
	"fmt"
	"github.com/Kirov7/FayGateway/dao"
	"github.com/Kirov7/FayGateway/middleware"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

func HTTPUriRewriteMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		service, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}
		serviceDetail := service.(*dao.ServiceDetail)

		for _, item := range strings.Split(serviceDetail.HTTPRule.UrlRewrite, ",") {
			items := strings.Split(item, " ")
			if len(items) != 2 {
				continue
			}
			regexp, err := regexp.Compile(items[0])
			if err != nil {
				fmt.Println("regexp.Compile err", err)
				continue
			}
			replacePath := regexp.ReplaceAll([]byte(c.Request.URL.Path), []byte(items[1]))
			c.Request.URL.Path = string(replacePath)
		}
		c.Next()
	}
}
