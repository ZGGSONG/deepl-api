package router

import (
	"deepl_api/api"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.GET("/", api.TranslateGet)
	r.GET("/translate", api.TranslateGet)
	r.POST("/translate", api.Translate)
	return r
}
