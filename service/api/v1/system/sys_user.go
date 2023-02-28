package system

import (
	"errors"

	"github.com/Gocyber-world/navigator-demo/model/common/response"
	"gorm.io/gorm"

	"github.com/Gocyber-world/navigator-demo/global"
	"github.com/Gocyber-world/navigator-demo/logger"
	systemReq "github.com/Gocyber-world/navigator-demo/model/system/request"
	systemResp "github.com/Gocyber-world/navigator-demo/model/system/response"
	"github.com/Gocyber-world/navigator-demo/service/system"
	"github.com/gin-gonic/gin"
)

// @Tags User
// @Summary 用户自行注册
// @accept application/json
// @Param data body systemReq.RegisterUser true " "
// @Param GOC-CAPTCHA header string true "recaptcha token"
// @Success 200 {object} response.Response "非生产环境中 test-mail-confirmation-uid 与 test-mail-confirmation-token header代表了用于激活的userId 和 token"
// @Router /v1/user/register [post]
func (b *BaseApi) RegisterUser(c *gin.Context) {
	var req systemReq.RegisterUser
	// 校验请求内容是否合法(密码长度、邮箱格式以及必要字段是否都带上了)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logger.Error(err.Error())
		response.FailWithMessage("Failed to parse request", c)
		return
	}

	// 判断邮箱是否已被使用
	if _, err := system.UserServiceApp.GetUserByEmail(req.Email); err == nil {
		logger.Errorf("email %s is occupied", req.Email)
		response.FailWithMessage("Email is occupied", c)
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Error(err.Error())
		response.FailWithMessage("Registration failed", c)
		return
	}

	// mysql中创建用户
	_, err = system.UserServiceApp.RegisterUser(req.Email, req.Password, req.NickName)
	if err != nil {
		logger.Error(err.Error())
		response.FailWithMessage("Registration failed", c)
		return
	}
	response.OkWithMessage("Success", c)
}

// @Tags User
// @Summary 通过邮箱和密码自行登录
// @accept application/json
// @Param data body systemReq.LoginUser true " "
// @Param GOC-CAPTCHA header string true "recaptcha token"
// @Success 200 {object} response.Response{data=systemResp.UserAccountResponse} "{"code":0,"data":"","msg":""}"
// @Router /v1/user/login [post]
func (b *BaseApi) LoginUser(c *gin.Context) {
	var req systemReq.LoginUser
	// 校验请求内容是否合法(密码长度、邮箱格式以及必要字段是否都带上了
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logger.Error(err.Error())
		response.FailWithMessage("Failed to parse request", c)
		return
	}

	user, err := system.UserServiceApp.LoginUser(req.Email, req.Password)
	if err != nil {
		logger.Error(err.Error())
		response.FailWithMessage("Login failure", c)
		return
	}

	token, user, err := system.AuthServiceApp.GenerateJWTWhenLogin(user.ID)
	if err != nil {
		logger.Error(err.Error())
		response.FailWithMessage("Login failure", c)
		return
	}

	c.SetCookie("goc-jwt", token, global.JWT_AUTH.Expiretime, "/", global.JWT_COOKIES_DOMAIN, false, false)
	// 返回账户上下文信息

	response.OkWithData(systemResp.UserAccountResponse{
		UserID: global.OBFUSE.Obfuscate(user.ID),
		Name:   user.NickName,
	}, c)
	logger.Infof("user %s login", user.NickName)
}
