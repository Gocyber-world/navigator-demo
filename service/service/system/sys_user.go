package system

import (
	"github.com/Gocyber-world/navigator-demo/global"
	"github.com/Gocyber-world/navigator-demo/logger"
	model "github.com/Gocyber-world/navigator-demo/model/system"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct{}

var UserServiceApp = new(UserService)

func (userService *UserService) GetUserByID(uid uint) (*model.SysUser, error) {
	user := &model.SysUser{
		Model: gorm.Model{
			ID: uid,
		},
	}

	if result := global.GVA_DB.Where(user).Take(user); result.Error != nil {
		return nil, result.Error
	} else {
		return user, nil
	}
}

func (us *UserService) GetUserByEmail(email string) (*model.SysUser, error) {
	user := &model.SysUser{
		Email: email,
	}

	if err := global.GVA_DB.Where(user).Take(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) CreateUser(u *model.SysUser) error {
	return global.GVA_DB.Create(u).Error
}

func (us *UserService) RegisterUser(email string, password string, nick string) (*model.SysUser, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	var newUser = &model.SysUser{
		NickName:              nick,
		Email:                 email,
		HashedPassword:        string(hashedPassword),
		AvatarModelUrl:        "",
		ProfilePicUrl:         "",
		BuiltopiaClientUserId: "",
	}

	if err = global.GVA_DB.Create(newUser).Error; err != nil {
		return nil, err
	}

	return newUser, err
}

func (us *UserService) LoginUser(email string, password string) (*model.SysUser, error) {
	u, err := us.GetUserByEmail(email)
	if err != nil {
		// 不区分用户不存在与用户密码错误
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password)); err != nil {
		return nil, err
	}

	return u, nil
}
