package main

import (
	"deepl_api/model"
	"deepl_api/util"
	"fmt"
	"github.com/goccy/go-json"
	"io"
	"log"
	"net/http"
	"strconv"
)

var (
	PORT = 8000
	URL  = "https://www2.deepl.com/jsonrpc"
)

func main() {
	// 注册请求处理函数
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/translate", translateHandler)

	// 启动HTTP服务器
	fmt.Printf("deepl server starte at %v...\n", PORT)

	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", PORT), nil); err != nil {
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

	var req model.Request
	var deeplResp model.DeepLResponse
	if err = json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Json Unmarshal Failed", http.StatusInternalServerError)
		return
	}

	// 开启 goroutine 发送 POST 请求并获取响应
	respChan := make(chan http.Response)
	go func() {
		text := req.Text
		sourceLang := req.SourceLang
		targetLang := req.TargetLang
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

		resp, err := util.HttpPost(URL, reqStr, headers)
		if err != nil {
			log.Printf("send to deepl post error: %v", err)
		}
		respChan <- resp
	}()

	_resp := <-respChan
	defer _resp.Body.Close()

	b, _ := io.ReadAll(_resp.Body)
	_ = json.Unmarshal(b, &deeplResp)
	if _resp.StatusCode == http.StatusOK {
		bytes, _ := json.Marshal(model.Response{
			Code: http.StatusOK,
			Data: deeplResp.Result.Texts[0].Text,
		})
		w.Write(bytes)

	} else {
		bytes, _ := json.Marshal(model.Response{
			Code: http.StatusTooManyRequests,
			Data: deeplResp.Error.Message,
		})
		w.Write(bytes)
	}

}
