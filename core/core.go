package core

import (
	"deepl_api/global"
	"deepl_api/model"
	"deepl_api/util"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func HttpCore(port int) {
	// 注册请求处理函数
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/translate", translateHandler)

	if port == 0 {
		port = 8000
	}

	// 启动HTTP服务器
	fmt.Printf("deepl server starte at %v...\n", port)

	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), nil); err != nil {
		log.Fatalf("deepl server start error: %v", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// 处理主页请求
	fmt.Fprintf(w, "POST {\"text\": \"input your content\", \"source_lang\": \"auto\", \"target_lang\": \"ZH\"} to /translate\n\n\ngithub.com/zggsong/stranslate\n")
}

func translateHandler(w http.ResponseWriter, r *http.Request) {
	// 定义返回值类型
	w.Header().Set("Content-Type", "application/json")
	// 处理翻译请求
	if r.Method == http.MethodGet {
		homeHandler(w, r)
		return
	} else if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 从请求体中获取需要翻译的文本
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad DeepLRequest", http.StatusBadRequest)
		return
	}
	//text := string(body)

	var req model.Request
	var deeplResp model.DeepLResponse
	if err = json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Json Unmarshal Failed", http.StatusInternalServerError)
		return
	}

	global.GLO_REQ_CH <- []string{req.Text, req.SourceLang, req.TargetLang}

	select {
	case _resp := <-global.GLO_RESP_CH:
		defer _resp.Body.Close()

		body, _ := io.ReadAll(_resp.Body)
		_ = json.Unmarshal(body, &deeplResp)
		if _resp.StatusCode == http.StatusOK {
			bytes, _ := json.Marshal(model.Response{
				Code: http.StatusOK,
				Data: deeplResp.Result.Texts[0].Text,
			})
			_, _ = w.Write(bytes)

		} else {
			bytes, _ := json.Marshal(model.Response{
				Code: http.StatusTooManyRequests,
				Data: deeplResp.Error.Message,
			})
			_, _ = w.Write(bytes)
		}
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
