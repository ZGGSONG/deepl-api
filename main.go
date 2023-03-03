package main

import (
	"deepl_api/core"
	"deepl_api/global"
	"net/http"
)

func main() {
	global.GLO_REQ_CH = make(chan []string)
	global.GLO_RESP_CH = make(chan http.Response)

	// 初始化服务器
	go core.HttpCore(8000)

	for {
		select {
		case sourceMsg := <-global.GLO_REQ_CH:
			go core.Handle(sourceMsg)
		}
	}
}
