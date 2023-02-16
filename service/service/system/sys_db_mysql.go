package system

import (
	"fmt"

	"github.com/Gocyber-world/navigator-demo/config"
	"github.com/Gocyber-world/navigator-demo/global"
	"github.com/Gocyber-world/navigator-demo/logger"
	model "github.com/Gocyber-world/navigator-demo/model/system"
	"github.com/Gocyber-world/navigator-demo/source/system"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// initMsqlDB 创建数据库并初始化 mysql
func (initDBService *InitDBService) InitMsqlDB(mysqlConfig config.Mysql) error {
	dsn := mysqlConfig.EmptyDsn()
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", mysqlConfig.Dbname)
	if err := initDBService.createDatabase(dsn, "mysql", createSql); err != nil {
		return err
	}
	if err := initDBService.ConnectGlobalDB(mysqlConfig); err != nil {
		global.GVA_DB = nil
		return err
	}
	if err := initDBService.initTables(); err != nil {
		global.GVA_DB = nil
		return err
	}
	if err := initDBService.initMysqlData(); err != nil {
		global.GVA_DB = nil
		return err
	}
	return nil
}

// 初始化全局数据库变量
func (initDBService *InitDBService) ConnectGlobalDB(mysqlConfig config.Mysql) error {
	logger.Info("connecting mysql database")
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       mysqlConfig.Dsn(),
		DefaultStringSize:         120,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	} else {
		global.GVA_DB = db
		return nil
	}
}

// initData mysql 初始化数据
func (initDBService *InitDBService) initMysqlData() error {
	return model.MysqlDataInitialize(
		system.User,
	)
}
