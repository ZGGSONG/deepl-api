package core

import (
	"api4Deeplx/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

// GinServe
//
//	@Description: 创建http服务
//	@param port 监听端口
func GinServe(port int) {
	r := gin.Default()
	r = router.CollectRoute(r)

	if port == 0 {
		r.Run(":8000")
	} else {
		r.Run(fmt.Sprintf(":%v", port))
	}
}
