package router

import (
	"demo/src/controller"
	"demo/src/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateServer() {
	// create router engine
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(middleware.CROSMiddleware(), middleware.AUTHMiddleware())

	// admin
	admin := router.Group("/admin")
	// login page
	admin.GET("/login", controller.AdminLoginGet)
	admin.POST("/login", controller.AdminLoginPost)
	// manage page
	admin.GET("/index/manage", controller.AdminIndexManagerGet)
	admin.POST("/index/manage", controller.AdminIndexManagerPost)
	// charge page
	admin.GET("/index/charge", controller.AdminIndexChargeGet)
	// report page
	admin.GET("/index/report", controller.AdminIndexReportGet)

	// client
	client := router.Group("/client")
	// login page
	client.GET("/login", controller.ClientLoginGet)
	client.POST("/login", controller.ClientLoginPost)
	// register page
	client.GET("/register", controller.ClientRegisterGet)
	client.POST("/register", controller.ClientRegisterPost)
	// information page
	client.GET("/index/information", controller.ClientIndexInformationGet)
	client.POST("/index/information", controller.ClientIndexInformationPost)
	// order page
	client.GET("/index/order", controller.ClientIndexOrderGet)
	client.POST("/index/order", controller.ClientIndexOrderPost)
	client.PATCH("/index/order", controller.ClientIndexOrderPatch)
	client.DELETE("/index/order", controller.ClientIndexOrderDelete)
	// car page
	client.GET("/index/car", controller.ClientIndexCarGet)
	client.POST("/index/car", controller.ClientIndexCarPost)
	client.PATCH("/index/car", controller.ClientIndexCarPatch)
	client.DELETE("/index/car", controller.ClientIndexCarDelete)

	// port
	http.ListenAndServe(":8848", router)
}
