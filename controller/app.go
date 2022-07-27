package controller

import "github.com/gin-gonic/gin"

type APPController struct{}

func APPRegister(router *gin.RouterGroup) {
	admin := APPController{}
	router.GET("/app_list", admin.APPList)
	router.GET("/app_detail", admin.APPDetail)
	router.GET("/app_stat", admin.APPStatistics)
	router.GET("/app_delete", admin.APPDelete)
	router.POST("/app_add", admin.APPAdd)
	router.POST("/app_update", admin.APPUpdate)
}

// APPList godoc
// @Summary 租户列表
// @Description 租户列表
// @Tags 租户管理
// @ID /app/app_list
// @Accept  json
// @Produce  json
// @Param info query string false "关键词"
// @Param page_size query int true "每页个数"
// @Param page_no query int true "当前页数"
// @Success 200 {object} middleware.Response{data=dto.ServiceListOutput} "success"
// @Router /app/app_list [get]
func (c *APPController) APPList(context *gin.Context) {

}

// APPDetail godoc
// @Summary 租户详情
// @Description 租户详情
// @Tags 租户管理
// @ID /service/service_delete
// @Accept  json
// @Produce  json
// @Param id query string true "服务ID"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_delete [get]
func (c *APPController) APPDetail(context *gin.Context) {

}

func (c *APPController) APPStatistics(context *gin.Context) {

}

func (c *APPController) APPDelete(context *gin.Context) {

}

func (c *APPController) APPAdd(context *gin.Context) {

}

func (c APPController) APPUpdate(context *gin.Context) {

}
