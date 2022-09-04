package core

import (
	"api4Deeplx/router"
	"github.com/gin-gonic/gin"
)

func GinServe() {
	r := gin.Default()
	r = router.CollectRoute(r)

	err := r.Run()
	if err != nil {
		return
	}
}
