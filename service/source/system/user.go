package system

import (
	"time"

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
	cTime := time.Now()
	defaultPasswordHash, err := bcrypt.GenerateFromPassword([]byte("Go121314"), bcrypt.DefaultCost)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if errors.Is(global.GVA_DB.Where("id = ?", 1006).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		if err := global.GVA_DB.Create(&[]system.SysUser{
			{Model: gorm.Model{ID: 1006}, NickName: "海格", HashedPassword: string(defaultPasswordHash), Email: "hagrid@gocyber.world", ConfirmAt: &cTime},
		}).Error; err != nil {
			return errors.Wrap(err, u.TableName()+" init data failed")
		}
	}
	if errors.Is(global.GVA_DB.Where("id = ?", 1005).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		if err := global.GVA_DB.Create(&[]system.SysUser{
			{Model: gorm.Model{ID: 1003}, NickName: "罗恩", HashedPassword: string(defaultPasswordHash), Email: "ron@gocyber.world", ConfirmAt: &cTime},
			{Model: gorm.Model{ID: 1004}, NickName: "麦格教授", HashedPassword: string(defaultPasswordHash), Email: "mg@gocyber.world", ConfirmAt: &cTime},
			{Model: gorm.Model{ID: 1005}, NickName: "斯内普", HashedPassword: string(defaultPasswordHash), Email: "snape@gocyber.world", ConfirmAt: &cTime},
		}).Error; err != nil {
			return errors.Wrap(err, u.TableName()+" init data failed")
		}
	}
	if errors.Is(global.GVA_DB.Where("id = ?", 1002).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		if err := global.GVA_DB.Create(&[]system.SysUser{
			{Model: gorm.Model{ID: 1000}, NickName: "哈利波特", HashedPassword: string(defaultPasswordHash), Email: "harry@gocyber.world", ConfirmAt: &cTime},
			{Model: gorm.Model{ID: 1001}, NickName: "马尔福", HashedPassword: string(defaultPasswordHash), Email: "malfoy@gocyber.world", ConfirmAt: &cTime},
			{Model: gorm.Model{ID: 1002}, NickName: "邓布利多 霍格伍兹负责人", HashedPassword: string(defaultPasswordHash), Email: "albus@gocyber.world", ConfirmAt: &cTime},
		}).Error; err != nil {
			return errors.Wrap(err, u.TableName()+" init data failed")
		}
	}
	if errors.Is(global.GVA_DB.Where("id = ?", 2000).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		if err := global.GVA_DB.Create(&[]system.SysUser{
			{Model: gorm.Model{ID: 2000}, NickName: "初始数据边界", HashedPassword: string(defaultPasswordHash), Email: "init-data-bound@gocyber.world", ConfirmAt: &cTime},
		}).Error; err != nil {
			return errors.Wrap(err, u.TableName()+" init data failed")
		}
	}
	return nil
}

func (u *user) CheckDataExist() bool {
	return !errors.Is(global.GVA_DB.Where("id = ?", 1005).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound)
}
