package server

import (
	"api4Deeplx/core"
	"api4Deeplx/global"
	"api4Deeplx/model/deepl"
	"api4Deeplx/util"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

var (
	URL         = "https://www2.deepl.com/jsonrpc"
	Refer       = "https://www.deepl.com/"
	ContentType = "application/json"
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
	targetLang := sourceMsg[1]
	timeSpan := util.GenerateTimestamp(text)
	id := util.CreateId()

	var reqStr = util.GenerateRequestStr(text, targetLang, timeSpan, id)

	var headers = make(map[string]string)
	headers["Content-Type"] = ContentType
	headers["Referer"] = Refer
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

//TODO: 修复目标语言 neutral
/*
This code produces the following output.

zh-Hans Chinese (Simplified)                    : neutral
zh      Chinese                                 : neutral
zh-Hant Chinese (Traditional)                   : neutral
zh-CHS  Chinese (Simplified) Legacy             : neutral
zh-CHT  Chinese (Traditional) Legacy            : neutral

zh-TW   Chinese (Traditional, Taiwan)           : specific
zh-CN   Chinese (Simplified, PRC)               : specific
zh-HK   Chinese (Traditional, Hong Kong S.A.R.) : specific
zh-SG   Chinese (Simplified, Singapore)         : specific
zh-MO   Chinese (Traditional, Macao S.A.R.)     : specific

*/

//if (name == "zh-CHT")
//{
//return "zh-Hant";
//}
//if (name == "zh-CHS")
//{
//return "zh-Hans";
//}
