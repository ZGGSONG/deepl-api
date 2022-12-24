package core

import (
	"api4Deeplx/router"
	"github.com/gin-gonic/gin"
)

func GinServe() {
	r := gin.Default()
	r = router.CollectRoute(r)

	err := r.Run(":8000")
	if err != nil {
		return
	}
}
