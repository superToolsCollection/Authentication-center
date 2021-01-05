package routers

import (
	"authentication-center/global"
	"authentication-center/internal/dao"
	"authentication-center/internal/middleware"
	"authentication-center/internal/routers/api"
	"authentication-center/internal/routers/user"
	"authentication-center/internal/service"
	"authentication-center/pkg/limiter"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

/**
* @Author: super
* @Date: 2021-01-04 20:57
* @Description:
**/

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	//r.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"https://foo.com"},
	//	AllowMethods:     []string{"PUT", "PATCH"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return origin == "https://github.com"
	//	},
	//	MaxAge: 12 * time.Hour,
	//}))
	// allow all origins解决跨域问题
	r.Use(cors.Default())
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.AppInfo())
	r.Use(middleware.Tracing())
	r.Use(middleware.RateLimiter(methodLimiters))
	//放到需要token的请求中
	//r.Use(middleware.JWT())
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//获取token
	authManager := dao.NewAuthManager("auth", global.DBEngine)
	authService := service.NewAuthService(authManager)
	authController := api.NewAuthController(authService)
	r.GET("/auth", authController.GetAuth)

	userManager := dao.NewUserManager("users", global.DBEngine)
	userService := service.NewUserService(userManager)
	userController := user.NewUserController(userService)

	userGroup := r.Group("/user")
	{
		userGroup.POST("/signin", userController.SignIn)
		userGroup.POST("/register", userController.Register)
		userGroup.PUT("/update", userController.Update)
	}
	return r
}
