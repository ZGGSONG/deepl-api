package core

import (
	"deeplapi/global"
	"deeplapi/model/deepl"
	"deeplapi/router"
	"deeplapi/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strconv"
)

// GinServe
//
//	@Description: 创建http服务
//	@param port 监听端口
func GinServe(port int) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r = router.CollectRoute(r)

	fmt.Println("starting deepl server at 8000...")
	if port == 0 {
		r.Run(":8000")
	} else {
		r.Run(fmt.Sprintf(":%v", port))
	}
}

var (
	URL         = "https://www2.deepl.com/jsonrpc"
	Refer       = "https://www.deepl.com/"
	ContentType = "application/json"
	UserAgent   = "DeepL-iOS/2.4.0 iOS 15.7.1 (iPhone14,2)"
)

// Handle
//
//	@Description: 翻译处理函数
//	@param sourceMsg
func Handle(sourceMsg []string) {
	var serveResp []string
	text := sourceMsg[0]
	sourceLang := sourceMsg[1]
	targetLang := sourceMsg[2]
	timeSpan := util.GenerateTimestamp(text)
	id := util.CreateId()
	sourceLang, regionalVariant := util.ConvertRegionalNameAndSourceLang(sourceLang, targetLang)
	var reqStr = util.GenerateRequestStr(text, sourceLang, targetLang, regionalVariant, timeSpan, id)

	var headers = make(map[string]string)
	headers["Content-Type"] = ContentType
	headers["Referer"] = Refer
	headers["Content-Length"] = strconv.Itoa(len(reqStr))
	headers["User-Agent"] = UserAgent
	//log.Printf("generate source: %v", reqStr)
	body, err := util.HttpPost(URL, reqStr, headers)
	if err != nil {
		global.GLO_RESP_CH <- []string{err.Error(), "NOT NULL"}
		return
	}
	//log.Printf("translate source: %v", string(body))
	var resp deepl.Response
	_ = json.Unmarshal(body, &resp)

	if resp.Result.Texts != nil {
		serveResp = []string{resp.Result.Texts[0].Text, ""}
		//log.Printf("translateText: %v\n", resp.Result.Texts[0].Text)
	} else {
		serveResp = []string{resp.Error.Message, string(rune(resp.Error.Code))}
		log.Printf("翻译出错")
		log.Printf("请求: %v", reqStr)
		log.Printf("返回: %v", string(body))
		//log.Printf("msg: %v\n", resp.Error.Message) //To many requests
		//log.Printf("code: %v\n", resp.Error.Code)   //1042911
	}
	global.GLO_RESP_CH <- serveResp
}
