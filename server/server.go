package server

import (
	"api4Deeplx/core"
	"api4Deeplx/global"
	"api4Deeplx/model/deepl"
	"api4Deeplx/util"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var (
	idNext int64
	URL    = "https://www2.deepl.com/jsonrpc"
)

func init() {
	idNext = rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000000000)
}

func Run() {
	//mgs channel
	global.GLO_REQ_CH = make(chan []string)
	global.GLO_RESP_CH = make(chan []string)

	//gin
	go core.GinServe()

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
	id := idNext + 1
	method := util.GenerateMethod(id)
	var reqStr = fmt.Sprintf("{\"jsonrpc\":\"2.0\",%vLMT_handle_texts\",\"params\":"+
		"{\"texts\":[{\"text\":\"%v\"}],"+
		"\"lang\":{\"target_lang\":\"%v\",\"source_lang_user_selected\":\"auto\"},"+
		"\"timestamp\":%v},\"id\":%v}", method, text, targetLang, timeSpan, id)

	var req = strings.NewReader(reqStr)
	respBytes, err := http.Post(URL, "application/json", req)
	if err != nil {
		log.Printf("请求出错: %v\n", err)
	}
	defer respBytes.Body.Close()
	body, _ := io.ReadAll(respBytes.Body)

	var resp deepl.Response
	_ = json.Unmarshal(body, &resp)

	if resp.Result.Texts != nil {
		serveResp = []string{resp.Result.Texts[0].Text, ""}
		log.Printf("translateText: %v\n", resp.Result.Texts[0].Text)
	} else {
		serveResp = []string{resp.Error.Message, string(rune(resp.Error.Code))}
		log.Printf("msg: %v\n", resp.Error.Message)
		log.Printf("code: %v\n", resp.Error.Code)
	}
	global.GLO_RESP_CH <- serveResp
}
