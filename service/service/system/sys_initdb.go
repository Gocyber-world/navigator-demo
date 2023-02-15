package system

import (
	"database/sql"

	"github.com/Gocyber-world/navigator-demo/global"
	"github.com/Gocyber-world/navigator-demo/logger"
	"github.com/Gocyber-world/navigator-demo/model/system"
)

type InitDBService struct{}

// initTables 初始化表
func (initDBService *InitDBService) initTables() error {
	return global.GVA_DB.AutoMigrate(
		system.SysUser{},
	)
}

// createDatabase 创建数据库(mysql)
func (initDBService *InitDBService) createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}
