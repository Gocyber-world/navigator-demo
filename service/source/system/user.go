package system

import (
	"github.com/Gocyber-world/navigator-demo/global"
	"github.com/Gocyber-world/navigator-demo/logger"
	"github.com/Gocyber-world/navigator-demo/model/system"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var User = new(user)

type user struct{}

func (u *user) TableName() string {
	return "sys_user"
}

func (u *user) Initialize() error {
	defaultPasswordHash, err := bcrypt.GenerateFromPassword([]byte("Navifans"), bcrypt.DefaultCost)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if errors.Is(global.GVA_DB.Where("id = ?", 1000).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		if err := global.GVA_DB.Create(&[]system.SysUser{
			{Model: gorm.Model{ID: 1000}, NickName: "简单男孩", HashedPassword: string(defaultPasswordHash), Email: "simple@demo.gocyber.world"},
		}).Error; err != nil {
			return errors.Wrap(err, u.TableName()+" init data failed")
		}
	}
	if errors.Is(global.GVA_DB.Where("id = ?", 2000).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		if err := global.GVA_DB.Create(&[]system.SysUser{
			{Model: gorm.Model{ID: 2000}, NickName: "初始数据边界", HashedPassword: string(defaultPasswordHash), Email: "init-data-bound@demo.gocyber.world"},
		}).Error; err != nil {
			return errors.Wrap(err, u.TableName()+" init data failed")
		}
	}
	return nil
}

func (u *user) CheckDataExist() bool {
	return !errors.Is(global.GVA_DB.Where("id = ?", 1000).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound)
}
