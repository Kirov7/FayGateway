package dto

import (
	"github.com/Kirov7/FayGateway/public"
	"github.com/gin-gonic/gin"
	"time"
)

type AdminLoginInput struct {
	UserName string `json:"username" form:"username" comment:"姓名" example:"admin" validate:"required,valid_username"` //管理员用户名
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"`               //管理员密码
}

type AdminLoginOutput struct {
	Token string `json:"token" form:"token" comment:"token" example:"token" validate:""` //token
}

type AdminSessionInfo struct {
	ID        int       `json:"id"`
	UserName  string    `json:"user_name"`
	LoginTime time.Time `json:"login_time"`
}

// BindValidParam 参数校验
func (param *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
