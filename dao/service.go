package dao

type ServiceDetail struct {
	Info     ServiceInfo `json:"info" description:"基本信息"`
	UserName string      `json:"user_name" gorm:"column:user_name" description:"管理员用户名"`
	Salt     string      `json:"salt" gorm:"column:salt" description:"盐"`
	Password string      `json:"password" gorm:"column:password" description:"密码"`
}
