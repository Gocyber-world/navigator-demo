package system

import (
	v1 "github.com/Gocyber-world/navigator-demo/api/v1"
	"github.com/gin-gonic/gin"
)

type V1Router struct{}

func (s *V1Router) InitV1ResourceRouter(Router *gin.RouterGroup) {
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	userRouter := Router.Group("")
	{
		userRouter.POST("user/register", baseApi.RegisterUser)
		userRouter.POST("user/login", baseApi.LoginUser)
	}

	// 需要认证的路由
	//accountRouter := Router.Group("").Use(global.JWT_AUTH.JWTAuthMiddleware())
	//{
	//	userRouter.POST("user/login/disabled", baseApi.LoginUser)
	//}
}
