package system

import (
	"go.uber.org/zap"
)

type InitDBFunc interface {
	Init() (err error)
}

const (
	Mysql           = "mysql"
	Pgsql           = "pgsql"
	InitSuccess     = "\n[%v] --> 初始数据成功!\n"
	AuthorityMenu   = "\n[%v] --> %v 视图已存在!\n"
	InitDataExist   = "\n[%v] --> %v 表的初始数据已存在!\n"
	InitDataFailed  = "\n[%v] --> %v 表初始数据失败! \nerr: %+v\n"
	InitDataSuccess = "\n[%v] --> %v 表初始数据成功!\n"
)

type InitData interface {
	TableName() string
	Initialize() error
	CheckDataExist() bool
}

// MysqlDataInitialize Mysql 初始化接口使用封装
func MysqlDataInitialize(inits ...InitData) error {
	for i := 0; i < len(inits); i++ {
		if inits[i].CheckDataExist() {
			zap.S().Infof(InitDataExist, Mysql, inits[i].TableName())
			continue
		}

		if err := inits[i].Initialize(); err != nil {
			zap.S().Infof(InitDataFailed, Mysql, inits[i].TableName(), err)
			return err
		} else {
			zap.S().Infof(InitDataSuccess, Mysql, inits[i].TableName())
		}
	}
	zap.S().Infof(InitSuccess, Mysql)
	return nil
}
