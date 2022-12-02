package controller

import (
	"encoding/base64"
	"github.com/Kirov7/FayGateway/dao"
	"github.com/Kirov7/FayGateway/dto"
	"github.com/Kirov7/FayGateway/middleware"
	"github.com/Kirov7/FayGateway/public"
	"github.com/Kirov7/go-config/lib"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type OauthController struct {
}

func OauthRegister(group *gin.RouterGroup) {
	oauth := &OauthController{}
	group.POST("/tokens", oauth.Tokens)
}

// Tokens godoc
// @Summary 获取Token
// @Description 获取Token
// @Tags OAUTH
// @ID /oauth/tokens
// @Accept  json
// @Produce  json
// @Param polygon body dto.TokensInput true "body"
// @Success 200 {object} middleware.Response{data=dto.TokensOutput} "success"
// @Router /admin/admin_info [get]
func (oauth *OauthController) Tokens(c *gin.Context) {
	params := &dto.TokensInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	splits := strings.Split(c.GetHeader("Authorization"), " ")
	if len(splits) != 2 {
		middleware.ResponseError(c, 2001, errors.New("用户名或密码格式错误"))
	}

	appSecret, err := base64.StdEncoding.DecodeString(splits[1])
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	// 取出 app_id secret
	// 生成 app_list
	// 匹配 app_id
	// 基于jwt生成token
	// 生成output

	parts := strings.Split(string(appSecret), ":")
	if len(parts) != 2 {
		middleware.ResponseError(c, 2003, errors.New("用户名或密码格式错误"))
		return
	}

	appList := dao.AppManagerHandler.GetAppList()
	for _, appInfo := range appList {
		if appInfo.AppID == parts[0] && appInfo.Secret == parts[1] {
			claims := jwt.StandardClaims{
				Issuer:    parts[0],
				ExpiresAt: time.Now().Add(public.JwtExpires * time.Second).In(lib.TimeLocation).Unix(),
			}
			token, err := public.JwtEncode(claims)
			if err != nil {
				middleware.ResponseError(c, 2004, err)
				return
			}
			output := &dto.TokensOutput{
				ExpiresIn:   public.JwtExpires,
				TokenType:   "Bearer",
				AccessToken: token,
				Scope:       "read_write",
			}
			middleware.ResponseSuccess(c, output)
			return
		}
	}
	middleware.ResponseError(c, 2005, errors.New("未匹配到正确的APP信息"))
}
