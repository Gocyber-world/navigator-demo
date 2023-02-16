package system

import (
	"time"

	"github.com/Gocyber-world/navigator-demo/global"
	"github.com/Gocyber-world/navigator-demo/middleware"
	"github.com/Gocyber-world/navigator-demo/model/system"
	"github.com/golang-jwt/jwt"
)

type AuthService struct{}

var AuthServiceApp = new(AuthService)

// 登录用的jwt生成方法
func (a *AuthService) GenerateJWTWhenLogin(uid uint) (string, *system.SysUser, error) {
	// 先尝试根据uid搜索组织管理员account
	var claims *middleware.CustomClaims
	// 先找出uid关联的所有account
	user, err := UserServiceApp.GetUserByID(uid)
	if err != nil {
		return "", nil, err
	}

	claims, err = a.generateJWTClaims(user.ID)
	if err != nil {
		return "", nil, err
	}
	token, err := global.JWT_AUTH.CreateToken(*claims)
	return token, user, err
}

func (a *AuthService) generateJWTClaims(userID uint) (*middleware.CustomClaims, error) {
	user, err := UserServiceApp.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return &middleware.CustomClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: time.Now().Unix() + int64(global.JWT_EXPIRE_TIME),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "gocyber",
			Subject:   "gocyber",
		},
		UserID:  global.OBFUSE.Obfuscate(user.ID),
		Name:    user.NickName,
		Version: 0,
	}, nil
}
