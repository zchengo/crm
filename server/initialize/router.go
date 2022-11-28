package initialize

import (
	"crm/api"
	"crm/global"
	"crm/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Router() {
	engine := gin.Default()

	// 开启跨域
	engine.Use(middleware.Cors())

	route := engine.Group("/api")

	{	
		// 用户模块
		route.GET("/user/verifycode", api.NewUserApi().GetVerifyCode)
		route.GET("/user/info", api.NewUserApi().GetInfo)
		route.PUT("/user/mail", api.NewUserApi().UpdateMail)
		route.PUT("/user/buy", api.NewUserApi().Buy)
		route.POST("/user/login", api.NewUserApi().Login)
		route.POST("/user/register", api.NewUserApi().Register)
		route.POST("/user/pass", api.NewUserApi().ForgotPass)
		route.DELETE("/user/logout", api.NewUserApi().Logout)
		route.DELETE("/user/delete", api.NewUserApi().Delete)

		// Jwt中间件
		route.Use(middleware.JwtAuth())

		// 客户模块
		route.GET("/customer/list", api.NewCustomerApi().QueryList)
		route.GET("/customer/info", api.NewCustomerApi().QueryInfo)
		route.GET("/customer/option", api.NewCustomerApi().QueryOption)
		route.POST("/customer/create", api.NewCustomerApi().Create)
		route.PUT("/customer/update", api.NewCustomerApi().Update)
		route.DELETE("/customer/delete", api.NewCustomerApi().Delete)

		// 合同模块
		route.GET("/contract/list", api.NewContractApi().QueryList)
		route.GET("/contract/info", api.NewContractApi().QueryInfo)
		route.GET("/contract/plist", api.NewContractApi().QueryPlist)
		route.PUT("/contract/update", api.NewContractApi().Update)
		route.POST("/contract/create", api.NewContractApi().Create)
		route.DELETE("/contract/delete", api.NewContractApi().Delete)

		// 产品模块
		route.GET("/product/list", api.NewProductApi().QueryList)
		route.GET("/product/info", api.NewProductApi().QueryInfo)
		route.POST("/product/create", api.NewProductApi().Create)
		route.PUT("/product/update", api.NewProductApi().Update)
		route.DELETE("/product/delete", api.NewProductApi().Delete)

		// 仪表盘模块
		route.GET("/dashboard/sum", api.NewDashboardApi().Summary)
		
	}

	// 启动、监听端口
	_ = engine.Run(fmt.Sprintf(":%v", global.Config.Server.Port))
}
