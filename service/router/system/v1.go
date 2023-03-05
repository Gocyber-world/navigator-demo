package system

import (
	v1 "github.com/Gocyber-world/navigator-demo/api/v1"
	"github.com/Gocyber-world/navigator-demo/global"
	"github.com/gin-gonic/gin"
)

type V1Router struct{}

func (s *V1Router) InitV1ResourceRouter(Router *gin.RouterGroup) {
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	userRouter := Router.Group("")
	{
		userRouter.POST("user/register", baseApi.RegisterUser)
		userRouter.POST("user/login", baseApi.LoginUser)
		userRouter.POST("builopia/customer", global.JWT_AUTH.JWTAuthMiddleware(), baseApi.SysUserRegisterBuiltopiaCustomer)
		userRouter.PATCH("user/info", global.JWT_AUTH.JWTAuthMiddleware(), baseApi.UpdateUserProfile)
		userRouter.GET("user/info", global.JWT_AUTH.JWTAuthMiddleware(), baseApi.GetUserInfo)
	}

	// 需要认证的路由
	//accountRouter := Router.Group("").Use(global.JWT_AUTH.JWTAuthMiddleware())
	//{
	//	userRouter.POST("user/login/disabled", baseApi.LoginUser)
	//}
}
