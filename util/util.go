package util

import (
	"api4Deeplx/model/deepl"
	"bytes"
	"encoding/json"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var NextId int64

func init() {
	//按规则生成结果
	NextId = int64(math.Round(rand.New(rand.NewSource(time.Now().UnixNano())).Float64()*10000.0) * 10000)
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
	// 当前时间戳
	num := time.Now().UnixMilli()
	// 转小写
	//texts = strings.ToLower(texts)
	// i 计数
	var num2 int64
	for _, text := range texts {
		if string(text) == "i" {
			num2++
		}
	}
	if num2 == 0 {
		num2 = 1
	}
	return num - num%num2 + num2
}

func adjustJsonContent(sourceReq string, id int64) (targetReq string) {
	var method string
	if (id+3)%13 == 0 || (id+5)%29 == 0 {
		method = "\"method\" : \""
	} else {
		method = "\"method\": \""
	}
	targetReq = strings.Replace(sourceReq, "\"method\":\"", method, 1)
	return
}

func HttpPost(url, reqStr string, header map[string]string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqStr)))
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func ConvertRegionalNameAndSourceLang(sourceLang, targetLang string) (sourceLangRet, regionalVariant string) {
	if sourceLang == "auto" {
		sourceLangRet = ""
	} else {
		sourceLangRet = sourceLang
	}
	if targetLang == "EN" {
		regionalVariant = "en-US"
	} else {
		regionalVariant = ""
	}
	return
}

func GenerateRequestStr(text, sourceLang, targetLang, regionalVariant string, timeSpan, id int64) (reqStr string) {
	req := deepl.Request{
		Jsonrpc: "2.0",
		Method:  "LMT_handle_texts",
		Params: deepl.ReqParams{
			Texts: []deepl.ReqParamsTexts{
				{
					Text:                text,
					RequestAlternatives: 3,
				},
			},
			Splitting: "newlines",
			Lang: deepl.ReqParamsLang{
				SourceLangUserSelected: sourceLang,
				TargetLang:             targetLang,
			},
			Timestamp: timeSpan,
			CommonJobParams: deepl.ReqParamsCommonJobParams{
				WasSpoken:       false,
				RegionalVariant: regionalVariant,
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
