package http_proxy_router

import (
	"github.com/Kirov7/FayGateway/controller"
	"github.com/Kirov7/FayGateway/http_proxy_middleware"
	"github.com/Kirov7/FayGateway/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	// todo 在压测时可以选择 gin.new(), 可以少打一些日志, 优化io性能
	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	oauth := router.Group("/oauth")
	oauth.Use(middleware.TranslationMiddleware())
	{
		controller.OauthRegister(oauth)
	}

	router.Use(
		http_proxy_middleware.HTTPAccessModeMiddleware(),

		http_proxy_middleware.HTTPFlowCountMiddleware(),
		http_proxy_middleware.HTTPFlowLimitMiddleware(),

		http_proxy_middleware.HTTPJwtAuthTokenMiddleware(),
		http_proxy_middleware.HTTPJwtFlowCountMiddleware(),
		http_proxy_middleware.HTTPJwtFlowLimitMiddleware(),

		http_proxy_middleware.HTTPWhiteListMiddleware(),
		http_proxy_middleware.HTTPBlackListMiddleware(),

		http_proxy_middleware.HTTPHeaderTransferMiddleware(),
		http_proxy_middleware.HTTPStripURIMiddleware(),
		http_proxy_middleware.HTTPUriRewriteMiddleware(),

		http_proxy_middleware.HTTPReverseProxyMiddleWare(),
	)
	return router
}
