package utils

import (
	"errors"

	"github.com/Gocyber-world/navigator-demo/global"
	"github.com/Gocyber-world/navigator-demo/middleware"
	"github.com/Gocyber-world/navigator-demo/model/common/enum"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	AccountID uint
	Type      string
	UserID    uint
	OrgID     uint
	Name      string
}

func GetUserInfo(c *gin.Context) (*UserInfo, error) {
	if claims, exists := c.Get("claims"); !exists {
		return nil, errors.New("get claims failed")
	} else {
		if waitUse, ok := claims.(*middleware.CustomClaims); !ok {
			return nil, errors.New("claims assertion failed")
		} else {
			userUintID, err := global.OBFUSE.Deobfuscate(waitUse.UserID)
			if err != nil {
				return nil, errors.New("unparseable uid")
			}
			return &UserInfo{UserID: userUintID, Name: waitUse.Name}, nil
		}
	}
}

func (u *UserInfo) IsOrgMember() bool {
	return u.Type == enum.OrgMember
}

func (u *UserInfo) IsCustomer() bool {
	return u.Type == enum.OrgCustomer
}

func (u *UserInfo) IsPersonal() bool {
	return u.Type == enum.Personal
}
