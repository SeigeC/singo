package service

import (
	"singo/model"
	"singo/serializer"
	"singo/serializer/handler"

	"github.com/gin-contrib/sessions"
	"gorm.io/gorm"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// setSession 设置session
func (service *UserLoginService) setSession(c *handler.Context, user model.User) {
	s := sessions.Default(c.Context)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *handler.Context) (handler.ActionResponse, error) {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, serializer.ErrDatabase.New(err)
		}
		return nil, serializer.ErrParamsMsg("账号或密码错误")
	}

	if user.CheckPassword(service.Password) == false {
		return nil, serializer.ErrParamsMsg("账号或密码错误")
	}

	// 设置session
	service.setSession(c, user)

	return user, nil
}
