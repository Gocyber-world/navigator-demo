package global

import (
	"github.com/Gocyber-world/navigator-demo/config"
	"github.com/Gocyber-world/navigator-demo/middleware"
	"github.com/Gocyber-world/navigator-demo/utils/obfuse"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB             *gorm.DB
	GVA_VP             *viper.Viper
	HOST               string
	JWT_COOKIES_DOMAIN string

	JWT_EXPIRE_TIME int
	OBFUSE          *obfuse.Obfuscator
	JWT_AUTH        *middleware.JWT
	STAGE           string
	MYSQL_CONFIG    config.Mysql
)
