package router

import (
	"github.com/gin-gonic/gin"
	"sample/api/apis"
)


func StartService() {
	router := gin.New()

	v1 := router.Group("v1")
	v1.GET("/test01",apis.Test01)

	router.Run(":878")
}
