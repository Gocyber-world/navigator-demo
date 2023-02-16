package system

import (
	"time"

	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	NickName           string     `json:"nickName" gorm:"size:64;comment:用户昵称"`
	Email              string     `json:"email" gorm:"index;size:64"`
	HashedPassword     string     `json:"hashedPassword" gorm:"size:64"`
	ConfirmationToken  string     `json:"confirmationToken" gorm:"size:64"`
	ConfirmAt          *time.Time `json:"confirmAt" gorm:"comment:激活时间"` // 在5.7以上的mysql中，默认情况下时间不可为零值
	ConfirmationSentAt *time.Time `json:"confirmationSentAt" gorm:"comment:激活邮件发送时间"`
}
