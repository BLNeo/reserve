package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"reserve/middleware"
	"reserve/router/apis/bookctl"
	"reserve/router/apis/testctl"
)

// InitRouter initialize routing information
func InitRouter(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(middleware.CORSMiddleware(), middleware.JWT())
	testRoute(r.Group("/test"))
	bookRoute(r.Group("/book"))
	manageRoute(r.Group("/manage"))
	clientRoute(r.Group("/manage"))
}

func testRoute(rg *gin.RouterGroup) {
	rg.GET("", testctl.GetTest)
}

func bookRoute(rg *gin.RouterGroup) {
	rg.POST("/add", bookctl.AddBook)
	rg.GET("/list", bookctl.ListBook)
}

// 后台端接口
func manageRoute(rg *gin.RouterGroup) {

}

// 客户端接口
func clientRoute(rg *gin.RouterGroup) {

}
