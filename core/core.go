package core

import (
	"deepl_api/global"
	"deepl_api/router"
	"deepl_api/util"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GinServe
//
//	@Description: 创建http服务
//	@param port 监听端口
func GinServe(port int) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r = router.CollectRoute(r)

	if port == 0 {
		port = 8080
	}

	fmt.Printf("starting deepl server at %v...\n", port)

	if err := r.Run(fmt.Sprintf("0.0.0.0:%v", port)); err != nil {
		log.Fatalf("starting deepl server error: %v", err)
	}
}

// Handle
//
//	@Description: 翻译处理函数
//	@param sourceMsg
func Handle(sourceMsg []string) {
	url := "https://www2.deepl.com/jsonrpc"
	text := sourceMsg[0]
	sourceLang := sourceMsg[1]
	targetLang := sourceMsg[2]
	timeSpan := util.GenerateTimestamp(text)
	id := util.CreateId()
	var reqStr = util.GenerateRequestStr(text, sourceLang, targetLang, timeSpan, id)

	var headers = make(map[string]string)
	headers["Content-Length"] = strconv.Itoa(len(reqStr))
	headers["User-Agent"] = "DeepL-iOS/2.4.0 iOS 15.7.1 (iPhone14,2)"
	headers["Accept"] = "*/*"
	headers["x-app-os-name"] = "IOS"
	headers["x-app-os-version"] = "15.7.1"
	headers["Accept-Language"] = "en-US,en;q=0.9"
	headers["Accept-Encoding"] = "gzip,deflate,br"
	headers["Content-Type"] = "application/json"
	headers["x-app-device"] = "iPhone14,2"
	headers["x-app-build"] = "353"
	headers["x-app-version"] = "2.4"
	headers["Referer"] = "https://www.deepl.com/"
	headers["Connection"] = "keep-alive"

	resp, err := util.HttpPost(url, reqStr, headers)
	if err != nil {
		log.Printf("translate failed, error: %v", err)
	}
	global.GLO_RESP_CH <- resp
}
