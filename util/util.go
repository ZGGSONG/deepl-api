package util

import (
	"strings"
	"time"
)

//
// GenerateTimestamp
//  @Description: 生成特定时间戳
//  @param texts
//  @return int64
//
func GenerateTimestamp(texts string) int64 {
	// 当前时间戳
	ts := time.Now().UnixMilli()
	// 转小写
	texts = strings.ToLower(texts)
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

//
// GenerateMethod
//  @Description: 根据id生成新的方法
//  @param id
//  @return method
//
func GenerateMethod(id int64) (method string) {
	if (id+3)%13 == 0 || (id+5)%29 == 0 {
		method = "\"method\" : \""
	} else {
		method = "\"method\": \""
	}
	return
}
