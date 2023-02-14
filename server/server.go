package server

import (
	"deepl_api/core"
	"deepl_api/global"
)

func Run() {
	global.GLO_REQ_CH = make(chan []string)
	global.GLO_RESP_CH = make(chan []string)

	go core.GinServe(9801)

	for {
		select {
		case sourceMsg := <-global.GLO_REQ_CH:
			go core.Handle(sourceMsg)
		}
	}
}
