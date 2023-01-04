package server

import (
	"api4Deeplx/core"
	"api4Deeplx/global"
	"api4Deeplx/model/deepl"
	"api4Deeplx/util"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"strconv"
)

var (
	URL         = "https://www2.deepl.com/jsonrpc"
	Refer       = "https://www.deepl.com/"
	ContentType = "application/json"
	UserAgent   = "DeepL-iOS/2.4.0 iOS 15.7.1 (iPhone14,2)"
)

func Run() {
	global.GLO_REQ_CH = make(chan []string)
	global.GLO_RESP_CH = make(chan []string)

	go core.GinServe(8000)

	for {
		select {
		case sourceMsg := <-global.GLO_REQ_CH:
			handle(sourceMsg)
		}
	}
}

func handle(sourceMsg []string) {
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
	log.Printf("generate source: %v", reqStr)
	body, err := util.HttpPost(URL, reqStr, headers)
	if err != nil {
		global.GLO_RESP_CH <- []string{err.Error(), "NOT NULL"}
		return
	}
	log.Printf("translate source: %v", string(body))
	var resp deepl.Response
	_ = json.Unmarshal(body, &resp)

	if resp.Result.Texts != nil {
		serveResp = []string{resp.Result.Texts[0].Text, ""}
		log.Printf("translateText: %v\n", resp.Result.Texts[0].Text)
	} else {
		serveResp = []string{resp.Error.Message, string(rune(resp.Error.Code))}
		log.Printf("msg: %v\n", resp.Error.Message) //To many requests
		log.Printf("code: %v\n", resp.Error.Code)   //1042911
	}
	global.GLO_RESP_CH <- serveResp
}
