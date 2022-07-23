package dto

import (
	"github.com/e421083458/FayGateway/public"
	"github.com/gin-gonic/gin"
)

type ServiceListInput struct {
	Info     string `json:"info" form:"info" comment:"关键词" example:"" validate:""`                      //关键词
	PageNo   int    `json:"page_no" form:"page_no" comment:"页数" example:"1" validate:"required"`        //所在页数
	PageSize int    `json:"page_size" form:"page_size" comment:"每页条数" example:"20" validate:"required"` //每页条数

}

type ServiceDeleteInput struct {
	ID int64 `json:"id" form:"id" comment:"服务ID" example:"20" validate:"required"` //关键词

}

type ServiceListItemOutput struct {
	ID          int64  `json:"id" form:"id" `                     //id
	ServiceName string `json:"service_name" form:"service_name" ` //服务名称
	ServiceDesc string `json:"service_desc" form:"service_desc" ` //服务描述
	LoadType    int    `json:"load_type" form:"load_type" `       //服务类型
	ServiceAddr string `json:"service_addr" form:"service_addr" ` //服务地址
	Qps         int64  `json:"qps" form:"qps" `                   //qps
	Qpd         int64  `json:"qpd" form:"qpd" `                   //qpd
	TotalNode   int    `json:"total_node" form:"total_node" `     //节点数

}

type ServiceListOutput struct {
	Total int64                   `json:"total" form:"total" comment:"总数" ` //总数
	List  []ServiceListItemOutput `json:"list" form:"list" comment:"列表"`    //列表
}

// BindValidParam 参数校验
func (param *ServiceListInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

// BindValidParam 参数校验
func (param *ServiceDeleteInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
