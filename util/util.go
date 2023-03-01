// Package util
// @Description: DeepL Rules
package util

import (
	"bytes"
	"deepl_api/model/deepl"
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var NextId int64

func init() {
	//按规则生成结果
	rand.Seed(time.Now().Unix())
	num := rand.Int63n(99999) + 8300000
	NextId = num * 1000
}

func CreateId() int64 {
	var num = NextId
	NextId = num + 1
	return num
}

// GenerateTimestamp
//
//	@Description: 生成特定时间戳
//	@param texts
//	@return int64
func GenerateTimestamp(texts string) int64 {
	iCount := int64(strings.Count(texts, "i"))
	// 当前时间戳
	ts := time.Now().UnixMilli()
	if iCount != 0 {
		iCount = iCount + 1
		return ts - ts%iCount + iCount
	} else {
		return ts
	}
}

func adjustJsonContent(sourceReq string, id int64) (targetReq string) {
	var method string
	if (id+3)%13 == 0 || (id+5)%29 == 0 {
		method = "\"method\" : \""
	} else {
		method = "\"method\": \""
	}
	targetReq = strings.Replace(sourceReq, "\"method\":\"", method, -1)
	return
}

// GenerateRequestStr
//
//	@Description: 构造请求json
//	@param text
//	@param sourceLang
//	@param targetLang
//	@param regionalVariant
//	@param timeSpan
//	@param id
//	@return reqStr
func GenerateRequestStr(text, sourceLang, targetLang string, timeSpan, id int64) (reqStr string) {
	req := deepl.Request{
		Jsonrpc: "2.0",
		Method:  "LMT_handle_texts",
		Params: deepl.ReqParams{
			Texts: []deepl.ReqParamsTexts{
				{
					Text:                text,
					RequestAlternatives: 0,
				},
			},
			Splitting: "newlines",
			Lang: deepl.ReqParamsLang{
				SourceLangUserSelected: sourceLang,
				TargetLang:             targetLang,
			},
			Timestamp: timeSpan,
			CommonJobParams: deepl.ReqParamsCommonJobParams{
				WasSpoken:    false,
				TranscribeAS: "",
				//RegionalVariant: regionalVariant,
			},
		},
		Id: id,
	}

	marshal, _ := json.Marshal(req)
	reqStr = string(marshal)

	var count int
	if len(string(marshal)) > 300 {
		count = 0
	} else {
		count = 3
	}
	req.Params.Texts[0].RequestAlternatives = count
	marshal, _ = json.Marshal(req)
	reqStr = string(marshal)

	reqStr = adjustJsonContent(reqStr, id)
	return
}

// HttpPost
//
//	@Description: 发送Post请求
//	@param url
//	@param reqStr
//	@param header
//	@return []byte
//	@return error
func HttpPost(url, reqStr string, header map[string]string) (http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqStr)))
	if err != nil {
		return http.Response{StatusCode: http.StatusInternalServerError}, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return http.Response{StatusCode: http.StatusInternalServerError}, err
	}

	return *resp, nil
}
