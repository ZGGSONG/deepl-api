package router

import (
	"api4Deeplx/api"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.GET("/", api.Index)
	r.POST("/translate", api.Translate)
	return r
}
