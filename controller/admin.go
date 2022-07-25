package controller

import (
	"encoding/json"
	"fmt"
	"github.com/e421083458/FayGateway/dao"
	"github.com/e421083458/FayGateway/dto"
	"github.com/e421083458/FayGateway/middleware"
	"github.com/e421083458/FayGateway/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
}

func AdminRegister(group *gin.RouterGroup) {
	adminInfo := &AdminController{}
	group.GET("/admin_info", adminInfo.AdminInfo)
	group.POST("/change_pwd", adminInfo.ChangePwd)
}

// AdminInfo godoc
// @Summary 管理员信息
// @Description 管理员信息
// @Tags 管理员接口
// @ID /admin/admin_info
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.AdminInfoOutput} "success"
// @Router /admin/admin_info [get]
func (changePwd *AdminController) AdminInfo(c *gin.Context) {
	sess := sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), adminSessionInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	// 1. 读取sessionKey对应 json 转换为结构体
	// 2. 取出数据然后封装输出结构体

	out := &dto.AdminInfoOutput{
		ID:           adminSessionInfo.ID,
		Name:         adminSessionInfo.UserName,
		LoginTime:    adminSessionInfo.LoginTime,
		Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Introduction: "SUPER ADMINISTRATOR",
		Roles:        []string{"admin"},
	}
	middleware.ResponseSuccess(c, out)
}

// ChangePwd godoc
// @Summary 管理员信息
// @Description 更改密码
// @Tags 管理员接口
// @ID /admin/change_pwd
// @Accept  json
// @Produce  json
// @Param polygon body dto.ChangePwdInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin/change_pwd [post]
func (changePwd *AdminController) ChangePwd(c *gin.Context) {
	params := &dto.ChangePwdInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	// 1. session读取用户信息到结构体 sessInfo
	// 2. sessInfo.ID 读取数据库信息 adminInfo
	// 3. params.password + adminInfo.salt sha256 saltPassword
	// 4. saltPassword => adminInfo.password 执行数据保存

	//session 读取用户信息到结构体
	sess := sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessInfo := dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), &adminSessInfo); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	// 从数据库中读取adminInfo
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	adminInfo := &dao.Admin{}
	adminInfo, err = adminInfo.Find(c, tx, &dao.Admin{UserName: adminInfo.UserName})
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}

	//生成新加盐密码逻辑
	saltPassword := public.GenSaltPassword(adminInfo.Salt, params.Password)
	adminInfo.Password = saltPassword

	//执行保存逻辑
	if err := adminInfo.Save(c, tx); err != nil {
		middleware.ResponseError(c, 2004, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}
