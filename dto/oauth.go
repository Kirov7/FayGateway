package dto

import (
	"github.com/Kirov7/FayGateway/public"
	"github.com/gin-gonic/gin"
)

type TokensInput struct {
	GrantType string `json:"grant_type" form:"grant_type" comment:"授权类型" example:"client_credentials" validate:"required"` //授权类型
	Scope     string `json:"scope" form:"scope" comment:"权限范围" example:"read_write" validate:"required"`                   //权限范围
}

type TokensOutput struct {
	AccessToken string `json:"access_token" form:"access_token"` //access_token
	ExpiresIn   int    `json:"expires_in" form:"expires_in"`     //expires_in
	TokenType   string `json:"token_type" form:"token_type"`     //token_type
	Scope       string `json:"scope" form:"scope"`               //scope
}

// BindValidParam 参数校验
func (param *TokensInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
