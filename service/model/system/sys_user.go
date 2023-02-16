package system

import (
	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	NickName       string `json:"nickName" gorm:"size:64;comment:用户昵称"`
	Email          string `json:"email" gorm:"index;size:64"`
	HashedPassword string `json:"hashedPassword" gorm:"size:64"`
}
