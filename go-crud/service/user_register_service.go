package service

import (
	"encoding/json"
	"fmt"
	"go-crud/model"
	"go-crud/serializer"
	"go-crud/util"
	"os"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
	Ticket          string `form:"ticket" json:"ticket" binding:"required"`
	Randstr         string `form:"randstr" json:"randstr" binding:"required"`
}

type Captcha struct {
	Response   string `json:"response"`
	Evil_level string `json:"evil_level"`
	Err_msg    string `json:"err_msg"`
}

// Valid 验证表单
func (service *UserRegisterService) Valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Status: 40001,
			Msg:    "两次输入的密码不相同",
		}
	}

	count := 0
	model.DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Status: 40001,
			Msg:    "昵称被占用",
		}
	}

	count = 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Status: 40001,
			Msg:    "用户名已经注册",
		}
	}

	return nil
}

// Captcha  验证码验证
func (service *UserRegisterService) CaptchaValid() *serializer.Response {
	aid := 2014850464
	AppSecretKey := ""   // 这里要填写自己的哦.. 我才不会给你呢！
	Ticket := service.Ticket // 拿过来前端传的ticket 和 randstr
	Randstr := service.Randstr
	url := fmt.Sprintf("https://ssl.captcha.qq.com/ticket/verify"+
		"?aid=%d"+
		"&AppSecretKey=%s"+
		"&Ticket=%s"+
		"&Randstr=%s", aid, AppSecretKey, Ticket, Randstr) // 构造url

	var txResponse Captcha
	response, err := util.GetResponse(url) // 自己写一个可以获取网页源码的函数
	if err != nil {
		return &serializer.Response{
			Status: 50003,
			Msg:    "服务器验证码模块出现异常，请检查",
		}
	}
	if err = json.Unmarshal(response, &txResponse); err != nil {
		return &serializer.Response{
			Status: 50003,
			Msg:    "格式化json出现异常",
		}
	}

	if txResponse.Response != "1" {
		return &serializer.Response{
			Status: 40002,
			Msg:    "验证不通过",
		}
	}
	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() (model.User, *serializer.Response) {
	randomString := util.RandStringRunes(16)
	user := model.User{
		Nickname:   service.Nickname,
		UserName:   service.UserName,
		Status:     model.Active,
		LiveAddr:   os.Getenv("LIVEFLV_ADDR") + randomString,
		LiveStream: os.Getenv("LIVEOBS_ADDR") + randomString,
	}

	// 表单验证
	if err := service.Valid(); err != nil {
		return user, err
	}

	// 腾讯验证码验证
	if err := service.CaptchaValid(); err != nil {
		return user, err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return user, &serializer.Response{
			Status: 40002,
			Msg:    "密码加密失败",
		}
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return user, &serializer.Response{
			Status: 40002,
			Msg:    "注册失败",
		}
	}

	return user, nil
}
