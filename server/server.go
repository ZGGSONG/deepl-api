package server

import (
	"deepl_api/model"
	"deepl_api/util"
	"github.com/fatih/color"
	"github.com/goccy/go-json"
	"io"
	"net/http"
	"strconv"
	"time"
)

type APIServer struct {
	baseURL    string
	httpClient *http.Client

	workerPool chan struct{}
}

func NewAPIServer(baseURL string, maxWorkers int) *APIServer {
	return &APIServer{
		baseURL:    baseURL,
		httpClient: &http.Client{},

		workerPool: make(chan struct{}, maxWorkers),
	}
}

func (s *APIServer) sendRequest(text, sourceLang, targetLang string) (http.Response, error) {
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

	resp, err := util.HttpPost(s.baseURL, reqStr, headers)
	if err != nil {
		return http.Response{}, err
	}
	return resp, nil
}

func (s *APIServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 处理非翻译请求
	if r.Method != http.MethodPost {
		w.Write([]byte("POST {\"text\": \"input your content\", \"source_lang\": \"auto\", \"target_lang\": \"ZH\"} to /translate\n\n\ngithub.com/zggsong/stranslate\n"))
		return
	}
	w.Header().Set("Content-Type", "application/json")

	// 从请求体中获取需要翻译的文本
	body, _ := io.ReadAll(r.Body)

	var req model.Request
	if err := json.Unmarshal(body, &req); err != nil {
		w.Write(marshal(model.Response{
			Code: http.StatusBadRequest,
			Data: "Error requesting json characters",
		}))
		return
	}

	if req.Text == "" {
		w.Write(marshal(model.Response{
			Code: http.StatusBadRequest,
			Data: "Please Enter Texts",
		}))
		return
	}

	// 等待空闲的 worker
	s.workerPool <- struct{}{}
	defer func() { <-s.workerPool }()

	// 发送 POST 请求并获取响应
	_resp, err := s.sendRequest(req.Text, req.SourceLang, req.TargetLang)
	if err != nil {
		w.Write(marshal(model.Response{
			Code: http.StatusInternalServerError,
			Data: err.Error(),
		}))
		return
	}

	var deeplResp model.DeepLResponse
	b, _ := io.ReadAll(_resp.Body)
	if err = json.Unmarshal(b, &deeplResp); err != nil {
		w.Write(marshal(model.Response{
			Code: http.StatusInternalServerError,
			Data: err.Error(),
		}))
		return
	}
	var resp = model.Response{
		Code: _resp.StatusCode,
	}

	if _resp.StatusCode == http.StatusOK {
		resp.Data = deeplResp.Result.Texts[0].Text
	} else {
		resp.Data = deeplResp.Error.Message
	}
	w.Write(marshal(resp))
}

func marshal(content model.Response) []byte {
	if content.Code != http.StatusOK {
		color.Red("[%v] Request Err: %v", time.Now().Format("2006-01-02 15:04:05"), content.Data)
	}
	bytes, _ := json.Marshal(content)
	return bytes
}
