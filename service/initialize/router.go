package initialize

import (
	_ "github.com/Gocyber-world/navigator-demo/docs"
	"github.com/Gocyber-world/navigator-demo/global"
	"github.com/Gocyber-world/navigator-demo/middleware"
	"github.com/Gocyber-world/navigator-demo/model/common/response"
	"github.com/Gocyber-world/navigator-demo/router"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/zap"
)

func Routers() *gin.Engine {
	if global.STAGE != "prod" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(
		middleware.GetInfoFromHeader(),
		ginzap.Ginzap(zap.L(), "", true),
		ginzap.RecoveryWithZap(zap.L(), true),
		middleware.Cors(),
	)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	systemRouter := router.RouterGroupApp.System
	PublicGroup := r.Group("")
	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			response.OkWithMessage("I am Ok with version 0", c)
		})
	}

	v1Group := r.Group("v1")
	{
		systemRouter.InitV1ResourceRouter(v1Group)
	}
	return r
}
