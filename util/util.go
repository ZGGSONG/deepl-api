package util

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
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
	ts := time.Now().UnixMilli()
	// 转小写
	//texts = strings.ToLower(texts)
	// i 计数
	var iCount int64
	for _, text := range texts {
		if string(text) == "i" {
			iCount++
		}
	}
	if iCount == 0 {
		iCount = 1
	}
	ret := ts - ts%iCount + iCount
	return ret
}

// GenerateMethod
//
//	@Description: 根据id生成新的方法
//	@param id
//	@return method
func GenerateMethod(id int64) (method string) {
	if (id+3)%13 == 0 || (id+5)%29 == 0 {
		method = "\"method\" : \""
	} else {
		method = "\"method\": \""
	}
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

func GenerateRequestStr(method, text, targetLang string, timeSpan, id int64) (reqStr string) {
	if targetLang == "ZH" {
		reqStr = fmt.Sprintf("{\"jsonrpc\":\"2.0\",%vLMT_handle_texts\",\"params\":"+
			"{\"texts\":[{\"text\":\"%v\"}],"+
			"\"lang\":{\"target_lang\":\"%v\",\"source_lang_user_selected\":\"auto\"},"+
			"\"timestamp\":%v,\"regionalVariant\":\"zh-CN\"},\"id\":%v}", method, text, targetLang, timeSpan, id)
		return
	}
	//无需区域变量
	reqStr = fmt.Sprintf("{\"jsonrpc\":\"2.0\",%vLMT_handle_texts\",\"params\":"+
		"{\"texts\":[{\"text\":\"%v\"}],"+
		"\"lang\":{\"target_lang\":\"%v\",\"source_lang_user_selected\":\"auto\"},"+
		"\"timestamp\":%v},\"id\":%v}", method, text, targetLang, timeSpan, id)
	return
}
